package migration

import (
	"bytes"
	"context"
	"errors"
	"fiap-hf-src/src/base/interfaces"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

var commands = map[string]*exec.Cmd{
	"pwd-windows": exec.Command("cmd", "/C", "echo", "%cd%"),
	"pwd":         exec.Command("bash", "-c", "pwd"),
}

type MigrationDB interface {
	Migrate() error
}

type migrationDB struct {
	DB interfaces.SQLDatabase
}

func NewMigration(db interfaces.SQLDatabase) MigrationDB {
	return &migrationDB{DB: db}

}

func (m migrationDB) Migrate() error {
	ctx := context.Background()

	if err := m.DB.Connect(); err != nil {
		return err
	}

	defer m.DB.Close()

	pwdOut, err := runCommandLinuxDarwinOrWindows("pwd")

	if err != nil {
		return err
	}

	pwdOutFormatted := formatCommandOutputToStr(pwdOut)

	path := filepath.Join(pwdOutFormatted, "infrastructure", "db", "DML")

	// for unity tests
	if strings.Contains(path, "/migration/infrastructure/db/DML") {
		path = filepath.Join(pwdOutFormatted, "..", "infrastructure", "db", "DML")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	list, err := file.Readdir(0)
	if err != nil {
		return err
	}

	files := make([]string, 0)

	for _, f := range list {
		files = append(files, f.Name())
	}

	sort.Slice(files[:], func(i, j int) bool {
		return files[i] < files[j]
	})

	for i := range files {

		if strings.Contains(files[i], ".example.sql") {
			continue
		}

		q, err := os.ReadFile(filepath.Join(path, files[i]))

		if err != nil {
			return err
		}

		if _, err := m.DB.ExecContext(ctx, string(q)); err != nil {
			return err
		}
	}

	return nil
}

func runCommandLinuxDarwinOrWindows(cmdName string) ([]byte, error) {
	if strings.EqualFold("windows", runtime.GOOS) {
		cmdName = cmdName + "-windows"
	}

	if _, ok := commands[cmdName]; ok {
		return commands[cmdName].Output()
	}

	return nil, errors.New("command not found")
}

func formatCommandOutputToStr(cmdOutput []byte) string {
	cmdOutput = bytes.ReplaceAll(cmdOutput, []byte("\n"), []byte(" "))
	cmdStr := string(cmdOutput)
	return strings.TrimSpace(cmdStr)
}
