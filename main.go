package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Weather struct {
	Name string `json:"name`
	Main struct {
		Temp     float64 `json:"temp`
		Pressure float64 `json:"pressure`
		Humidity float64 `json:"humidity`
	}
	Wind struct {
		Speed float64 `json:"speed`
	}
	Sys struct {
		Country string  `json:"country`
		Sunrise float64 `json:"sunrise`
		Sunset  float64 `json:"sunset`
	}
}

var data Weather
var city string

func main() {
	fmt.Println("Weather Application")
	// initialising router
	router := mux.NewRouter()
	// handler to route request
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/{city}", WeatherData)
	http.ListenAndServe(":4000", router)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>" + "Weather Application" + "<h1>"))
}
func WeatherData(w http.ResponseWriter, r *http.Request) {
	// retreiving city name from request
	params := mux.Vars(r)
	city := params["city"]
	key := "b9dcd870965d814848774508704c8582"
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + key
	// fetching data from the given url
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// fetched data is in json format need to convert it into struct to process
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	// assigning fetched data into variables
	country := data.Sys.Country
	city = data.Name
	temp := strconv.FormatFloat(data.Main.Temp-270, 'f', -1, 64)
	pressure := strconv.FormatFloat(data.Main.Pressure, 'f', -1, 64)
	humidity := strconv.FormatFloat(data.Main.Humidity, 'f', -1, 64)
	speed := strconv.FormatFloat(data.Wind.Speed*3.6, 'f', -1, 64)
	// print statements for server
	w.Write([]byte("<h1>" + "Fetched Weather Data is " + "<h1>"))
	w.Write([]byte("<h1>" + "country is : " + country + "<h1>"))
	w.Write([]byte("<h1>" + "city is : " + city + "<h1>"))
	w.Write([]byte("<h1>" + "temp is : " + temp + " celsius" + "<h1>"))
	w.Write([]byte("<h1>" + "pressure is : " + pressure + " milibar" + "<h1>"))
	w.Write([]byte("<h1>" + "humidity is : " + humidity + " %" + "<h1>"))
	w.Write([]byte("<h1>" + "wind speed is : " + speed + " km/hr" + "<h1>"))
}
