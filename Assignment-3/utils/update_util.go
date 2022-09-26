package utils

func UpdateWeather(windValue, waterValue int8) string {
	update := ""
	if windValue <= 6 {
		temp := "aman"
		update = temp
	} else if windValue >= 7 && windValue <= 15 {
		temp := "siaga"
		update = temp
	} else if windValue > 15 {
		temp := "bahaya"
		update = temp
	} else if waterValue < 5 {
		temp := "aman"
		update = temp
	} else if waterValue >= 6 && waterValue <= 8 {
		temp := "siaga"
		update = temp
	} else if waterValue > 8 {
		temp := "bahaya"
		update = temp
	}
	return update
}
