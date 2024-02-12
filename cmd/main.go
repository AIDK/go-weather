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

	// create two channels to pass the locationRes and weather data
	locationRes := make(chan *[]models.LocationResponse)
	weatherRes := make(chan *models.WeatherResponse)
	// we get the geo location in order to pass it to the getWeather function
	// we use the <- operator to pass the locationRes channel to the getGeoLocation function
	// and get the location from the channel
	go getGeoLocation(locationRes)
	// get the location from the channel and pass it to the getWeather function
	// we use the <- operator to receive the value from the channel
	// and pass it to the getWeather function
	go getWeather(weatherRes, <-locationRes)
	// get the weather from the channel and print it
	weather := <-weatherRes

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

// getGeoLocation returns the location for a given location.
// The location is passed to the function via the out channel and the location is returned via the same channel.
// The out channel is closed when the function returns.
func getGeoLocation(out chan<- *[]models.LocationResponse) {

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

	// pass the location to the out channel and close the channel (we are done)
	out <- locationResponse
	close(out)
}

// getWeather returns the weather for a given location.
// The location is passed to the function via the out channel and the location is returned via the same channel.
// The out channel is closed when the function returns.
// The function takes a locationResponse and a weatherResponse as input.
func getWeather(out chan<- *models.WeatherResponse, locationResponse *[]models.LocationResponse) {

	// get the latitude and longitude from the locationResponse
	// we use the * operator to dereference the pointer and get the value
	latitude := (*locationResponse)[0].Lat
	longitude := (*locationResponse)[0].Lon

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

	// pass the weather to the out channel and close the channel (we are done)
	out <- weatherResponse
	close(out)
}
