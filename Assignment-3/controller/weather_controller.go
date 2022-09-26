package controller

import (
	"assignment3/constant"
	"assignment3/service"
	"assignment3/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type weatherController struct {
	r  *mux.Router
	ws service.WeatherService
}

func (w *weatherController) InitController() {
	routes := w.r.NewRoute().PathPrefix(constant.WEATHER_API_PATH).Subrouter()

	routes.HandleFunc("", w.updateValue()).Methods(http.MethodPut)
	routes.HandleFunc("/update", w.getUpdate()).Methods(http.MethodGet)
}

func ProvideWeatherController(r *mux.Router, ws service.WeatherService) *weatherController {
	return &weatherController{r: r, ws: ws}
}

func (w *weatherController) updateValue() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := w.ws.UpdateValue(r.Context())
		if err != nil {
			utils.NewErrorResponse(rw, err)
		}
		utils.NewSuccessResponsWriter(rw, http.StatusCreated, "Updated", nil)
	}
}

func (w *weatherController) getUpdate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		res, err := w.ws.GetUpdate(r.Context())
		if err != nil {
			utils.NewErrorResponse(rw, err)
		}

		utils.NewSuccessResponsWriter(rw, http.StatusOK, "SUCCESS", res)
	}
}
