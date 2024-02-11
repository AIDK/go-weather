package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	utils "github.com/aidk/go-weather/helpers"
	"github.com/aidk/go-weather/models"
	_ "github.com/joho/godotenv/autoload"
)

const baseUrl string = "https://api.openweathermap.org/"

func main() {

	location := getGeoLocation()
	weather := getWeather(location)

	fmt.Printf("%s (%s): %.0fC, %s\n",
		weather.City.Name,
		weather.City.Country,
		weather.List[0].Main.Temp/10,
		weather.List[0].Weather[0].Description)

	// print the next 5 forecasts (3-hourly)
	for _, w := range weather.List {

		// we need to format date because the API is returning a Unix timestamp
		date := time.Unix(int64(w.Dt), 0)
		// if we have a date in the past, we skip it (we only want future forecasts)
		if date.Before(time.Now()) {
			continue
		}

		// print the forecast (date, temperature, weather description)
		message := fmt.Sprintf("%s - %.0fC, %s (%s)\n",
			date.Format("Mon Jan 2 15:04"),
			w.Main.Temp/10,
			w.Weather[0].Main,
			w.Weather[0].Description)

		fmt.Print(message)
	}

}

// getGeoLocation returns the location for a given query
func getGeoLocation() *[]models.LocationResponse {

	// get the location from the command line arguments (if present)
	location := utils.Location()
	response, err := http.Get(
		baseUrl +
			"geo/1.0/direct" +
			"?q=" + location +
			"&limit=1" +
			"&appid=" + os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	// close the response body when the function returns
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("Error")
	}

	// read the response body into a byte slice
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal the JSON byte slice into a Go data structure
	// we defined the Location struct in the models package
	locationResponse := &[]models.LocationResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		panic(err)
	}

	return locationResponse
}

// getWeather returns the weather for a given location
func getWeather(location *[]models.LocationResponse) *models.WeatherResponse {

	// get the latitude and longitude from the location
	latitude := (*location)[0].Lat
	longitude := (*location)[0].Lon

	response, err := http.Get(
		baseUrl +
			"/data/2.5/forecast" +
			"?lat=" + fmt.Sprintf("%f", latitude) +
			"&lon=" + fmt.Sprintf("%f", longitude) +
			"&appid=" + os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	// close the response body when the function returns
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(err)
	}

	// read the response body into a byte slice
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal the JSON byte slice into a Go data structure
	weatherResponse := &models.WeatherResponse{}
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		panic(err)
	}

	return weatherResponse
}
