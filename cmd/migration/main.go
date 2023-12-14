package main

import (
	"bufio"
	"context"
	"fiap-hf-src/infrastructure/db/migration"
	"fiap-hf-src/pkg/postgres"
	"log"
	"os"
	"regexp"
)

var regexEnvs = regexp.MustCompile(`(\S+)=(\S+)`)

func init() {
	if err := defineEnvs(".env"); err != nil {
		log.Fatalf("Error to load .env -> %v", err)
	}
}

func main() {
	ctx := context.Background()

	db := postgres.NewPostgresDB(
		ctx,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	m := migration.NewMigration(db)

	if err := m.Migrate(); err != nil {
		log.Fatalf("error -> %v", err)
	}

	log.Printf("Migration runned with success")
}

func defineEnvs(filename string) error {
	file, err := os.Open(filename)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error on close -> %v", err)
		}
	}(file)

	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		envMatch := regexEnvs.FindStringSubmatch(sc.Text())
		if envMatch != nil {
			err := os.Setenv(envMatch[1], envMatch[2])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
