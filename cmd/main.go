package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lgustavopalmieri/ultrafactory-oee/configs"
	"github.com/lgustavopalmieri/ultrafactory-oee/internal/adapters/web/routes"
	"github.com/lgustavopalmieri/ultrafactory-oee/internal/adapters/web/server"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	// postgresdb.RunPostgresMigrations(cfg.MigrationURL, cfg.DBSource)
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Unable to connect to the database: %v", err))
	}
	fmt.Println("Successfully connected to the database.")

	webserver := server.NewWebServer(":4003")
	routes.SetupTestRoutes(webserver,db)
	webserver.Start()
}