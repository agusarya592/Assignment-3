package entity

type Weather struct {
	Wind  string `db:"wind"`
	Water string `db:"water"`
}

type StatusUpdate struct {
	Status  string `db:"report"`
	Weather Weather
}

type WeatherResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Wind  string `json:"wind"`
	Water string `json:"water"`
}

func CreateWeatherReport(s StatusUpdate) *WeatherResponse {
	return &WeatherResponse{
		Status: s.Status,
		Data: Data{
			Wind:  s.Weather.Wind,
			Water: s.Weather.Water,
		},
	}
}

func CreateWeatherResponse(s StatusUpdate) *WeatherResponse {
	res := *CreateWeatherReport(s)
	return &res
}
