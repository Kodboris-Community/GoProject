package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"kodboris/api"
	_ "kodboris/db/sqlc"
	db "kodboris/db/sqlc"
	"kodboris/util"
	"log"
	"path/filepath"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env: %w", err)
	}
	fmt.Printf("Database driver: %v\n", config.DbDriver)

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	// Fail startup, if connection to the db can not be established
	err = conn.Ping() // Check if the connection is valid
	if err != nil {
		log.Fatalf("Application can not start, Error connecting to db: %v", err) // Fatal log and terminate
	}

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Migration error %v\n", err)
	}

	// Pass in the absolute path to the migration directory
	absPath, err := filepath.Abs("./db/migration")
	if err != nil {
		log.Fatal("Error getting absolute path: ", err)
	}

	// Pass in the db parameters for migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+absPath, config.DbDriver, driver)
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}

	// call to migrate up method
	err = m.Up()
	//fmt.Printf("Database Migration is successful!!!!!!\n")
	if err != nil {
		fmt.Printf("%v to the migration\n", err)
		if err != migrate.ErrNoChange {
			log.Fatalf("Error applying migrations %v\n", err)
		}
	}

	// Release resources after db migration
	defer m.Close()

	//// Connection to rabbitmq
	//_, err = mq.ConnectToRabbitmq(config.RabbitMq)

	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	address := config.ServerAddress
	err = server.Start(address, store)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}

}
