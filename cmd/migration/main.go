package main

import (
	"bufio"
	"context"
	"fiap-hf-src/migration"
	"fiap-hf-src/src/external/db/rds/postgres"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	if err := defineEnvs(".env"); err != nil {
		log.Fatalf("Error to load .env -> %v", err)
	}
}

func main() {
	ctx := context.Background()

	dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	db := postgres.NewPostgresDB(
		ctx,
		os.Getenv("DB_REGION"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		dbDuration,
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
		indexComment := strings.Index(sc.Text(), "#")
		if indexComment != -1 && len(strings.TrimSpace(sc.Text()[:indexComment])) == 0 {
			continue
		}
		envEqualSign := strings.Index(sc.Text(), "=")
		if envEqualSign != -1 {
			envMatchKey := sc.Text()[:envEqualSign]
			envMatchValue := sc.Text()[envEqualSign+1:]
			if len(envMatchKey) != 0 || len(envMatchValue) != 0 {
				err := os.Setenv(envMatchKey, envMatchValue)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}
