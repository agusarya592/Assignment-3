package route

import (
	"assignment3/controller"
	"assignment3/service"
	"database/sql"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router, db *sql.DB) {
	webRouter := router.NewRoute().PathPrefix("/api").Subrouter()

	weatherService := service.ProvideWeatherService(db)
	weatherHandler := controller.ProvideWeatherController(webRouter, weatherService)
	weatherHandler.InitController()
}
