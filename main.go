package main

import (
	"database/sql"
	"log"

	"github.com/Llala/simplecat/api"
	db "github.com/Llala/simplecat/db/sqlc"
	"github.com/Llala/simplecat/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:")
	}

	// if config.Environment == "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ")
	}

	// runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	runGinServer(config, store)

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:")
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server:")
	}

}
