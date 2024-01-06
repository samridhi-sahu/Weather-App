# Golang Weather App

A simple Golang Weather App that fetches current weather data based on city names using the OpenWeather API.

## Features

- **Weather Data Retrieval**: Users can request current weather data for a specific city.
- **OpenWeather API Integration**: Utilizes the OpenWeather API to fetch real-time weather information.
- **Web Server Display**: The fetched weather data is displayed on a web server.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Golang installed on your machine**: Make sure you have the Go programming language installed.
- **OpenWeather API Key**: Obtain an API key by signing up on the [https://openweathermap.org/]

## To Run

1. **Clone the repository**:

   ```bash
   git clone https://github.com/samridhi-sahu/Weather-App.git
   
2. **Navigate to the project directory**:**:

   ```bash
   cd your-repo

3. **Install Dependencies**:

   ```bash
   go get -u github.com/gorilla/mux

4. Paste your particular openweather api key, key := "your key"
5. **Run**:

   ```bash
   go run main.go

## Usage

1. **Start Server**:

   ```bash
   http://localhost:4000

2. **Access the weather data by making a request to the following endpoint**:

   ```bash
   http://localhost:4000/city

3. **This will show fetched data on your local server**


