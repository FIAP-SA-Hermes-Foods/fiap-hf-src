package main

import (
	"bufio"
	"hermes-foods/infrastructure/db/migration"
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
	if err := migration.DBMigration(); err != nil {
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
