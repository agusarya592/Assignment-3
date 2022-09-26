package database

import (
	"assignment3/config"
	"database/sql"
	"fmt"
	"log"
)

func GetDb() *sql.DB {
	log.Printf("INFO GetDb => starting database connection")

	cfg := config.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Address, cfg.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("ERROR! GetDb sql open connection fatal error => %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ERROR! GetDb database ping fatal error => %v", err)
	}
	log.Printf("INFO GetDb database connection: estabilished successfully => %s\n", dsn)
	return db
}
