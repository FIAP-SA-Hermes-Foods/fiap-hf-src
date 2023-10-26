package migration

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func DBMigration() error {
	ctx := context.Background()

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", url)

	if err != nil {
		return err
	}

	conn, err := db.Conn(ctx)

	if err != nil {
		return err
	}

	defer conn.Close()

	// TODO: Run this with a PWD command checking by OS system
	q, err := os.ReadFile("./infrastructure/db/scripts/001_init_db.sql")

	if err != nil {
		return err
	}

	if _, err := conn.ExecContext(ctx, string(q)); err != nil {
		return err
	}

	return nil
}
