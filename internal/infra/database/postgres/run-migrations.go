package postgres

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunPostgresMigrations(db *sql.DB) {
	// Set the embedded migrations as the source
	goose.SetBaseFS(embedMigrations)
	
	// Set the dialect to Postgres
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("Cannot set database dialect: ", err)
	}
	
	// Run migrations
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	
	log.Println("Database migration successfully applied!")
}
