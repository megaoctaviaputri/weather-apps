package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"weather-apps/models"

	"github.com/labstack/echo"
)

func GetWeatherController(c echo.Context) error {
	name := c.Param("name")
	appid := "94b82340db3fe3d02225564c51932352"

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + name + "&appid=" + appid

	response, _ := http.Get(url)

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var weatherTemp models.WeatherData
	json.Unmarshal(responseData, &weatherTemp)

	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":        name,
		"temperature": weatherTemp.Main.Temp,
	})
}
