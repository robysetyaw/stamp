package stamp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func FetchOpenWeatherMap() ([]string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	apiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	baseURL := "https://api.openweathermap.org/data/2.5/forecast"

	params := url.Values{}
	params.Add("q", "Jakarta")
	params.Add("appid", apiKey)
	params.Add("units", "metric")
	params.Add("lang", "id")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	response, err := http.Get(fullURL)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var forecast ForecastResponse

	err = json.Unmarshal(body, &forecast)

	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	var results []string
	takenDates := map[string]bool{}
	for _, item := range forecast.List {
		t, err := time.Parse("2006-01-02 15:04:05", item.DtTxt)
		if err != nil {
			continue
		}
		dateKey := t.Format("2006-01-02")
		if takenDates[dateKey] {
			continue
		}
		takenDates[dateKey] = true
		date := t.Format("Mon, 02 Jan 2006")
		info := fmt.Sprintf("%s : %.1f°C", date, item.Main.Temp)
		results = append(results, info)
	}

	return results, nil
}

type ForecastResponse struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
}
