package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fayob/go_micro/auth_service/data"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main()  {
	
	// TODO connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres")
	}
	
	// set up config
	app := Config{
		DB: conn,
		Models: data.New(conn),
	}
	
	// define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	
	log.Println("Starting authentication service")
	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
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
	dns := os.Getenv("DNS")
	
	connection, err := openDB(dns)
	if err != nil {
		log.Println(err)
		log.Panic("Postgres not yet ready...")
	}
	log.Println("Connected to Postgres!")
	return connection
}