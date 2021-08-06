package openweathermapgo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	. "openweathermap-go"
	"os"
	"testing"
	"time"
)

var (
	mux                   *http.ServeMux
	server                *httptest.Server
	currentWeatherService CurrentWeatherService
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	currentWeatherService = NewCurrentWeatherService(&http.Client{Timeout: 10 * time.Second}, os.Getenv("OWM_API_KEY"), IMPERIAL, FRENCH, &CurrentWeatherServiceOptions{BaseURL: server.URL})

	return func() {
		server.Close()
	}
}

func TestCurrentWeatherByCityName(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/data/2.5/weather", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CurrentWeather{
			Main: Main{
				Temp: 10.8,
			},
		})
	})

	currentWeather, err := currentWeatherService.CurrentByCityName("Davis")

	if err != nil {
		t.Fatal(err)
	}

	if currentWeather.Main.Temp != 10.8 {
		t.Errorf("Expected 10.8 got %f", currentWeather.Main.Temp)
	}
}

func TestCurrentWeatherByZipCode(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/data/2.5/weather", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CurrentWeather{
			Main: Main{
				Temp: 10.8,
			},
		})
	})

	currentWeather, err := currentWeatherService.CurrentByZipCode("95618")
	if err != nil {
		t.Fatal(err)
	}

	if currentWeather.Main.Temp != 10.8 {
		t.Errorf("Expected 10.8 got %f", currentWeather.Main.Temp)
	}
}

func TestCurrentWeatherByGeoPos(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/data/2.5/weather", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CurrentWeather{
			Main: Main{
				Temp: 10.8,
			},
		})
	})

	currentWeather, err := currentWeatherService.CurrentByGeoPos(10.32, 132)
	if err != nil {
		t.Fatal(err)
	}

	if currentWeather.Main.Temp != 10.8 {
		t.Errorf("Expected 10.8 got %f", currentWeather.Main.Temp)
	}
}
