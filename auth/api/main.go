package main

import (
	"auth/data"
	"broker/cmd/util"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var port string = util.AssertEnv("SERVICE_PORT")
var dsn string = util.AssertEnv("DSN")

const DB_CONN_RETRIES = 15

type App struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service.")

	conn := connectToDB()
	if conn == nil {
		log.Panic("failed to connect to Postgres")
	}

	app := App{
		DB:     conn,
		Models: data.New(conn),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	for count := 0; count < DB_CONN_RETRIES; count++ {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")

		} else {
			log.Println("Connected to Postgres.")
			return connection
		}
	}

	return nil
}
