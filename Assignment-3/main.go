package main

import (
	"assignment3/config"
	"assignment3/database"
	"assignment3/route"
	"assignment3/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	config.Init()
	cfg := config.GetConfig()
	root := mux.NewRouter()
	db := database.GetDb()

	route.Init(root, db)
	s := server.ProvideServer(cfg.ServerAddress, root)
	s.ListenAndServe()
}
