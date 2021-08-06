package openweathermapgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const API_URL string = "https://api.openweathermap.org"
const CURRENT_WEATHER_ENDPOINT = "/data/2.5/weather?%s"

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CurrentWeatherService struct {
	client *http.Client
	apiURL string
	apiKey string
	unit   Unit
	lang   Language
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Sys struct {
	Type    int64   `json:"type"`
	Id      int64   `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

type CurrentWeather struct {
	GeoPos     Coordinates `json:"coord"`
	Weather    []Weather   `json:"weather"`
	Main       Main        `json:"main"`
	Visibility int64       `json:"visibility"`
	Base       string      `json:"base"`
	Wind       Wind        `json:"wind"`
	Clouds     Clouds      `json:"clouds"`
	Dt         int64       `json:"dt"`
	Sys        Sys         `json:"sys"`
	Timezone   int64       `json:"timezone"`
	ID         int64       `json:"id"`
	CityName   string      `json:"city_name"`
	Cod        int64       `json:"cod"`
}

type CurrentWeatherServiceOptions struct {
	BaseURL string
}

func NewCurrentWeatherService(c *http.Client,
	apiKey string,
	unit Unit,
	lang Language,
	options *CurrentWeatherServiceOptions) CurrentWeatherService {
	var apiURL string = API_URL + CURRENT_WEATHER_ENDPOINT
	if options != nil && options.BaseURL != "" {
		apiURL = options.BaseURL + CURRENT_WEATHER_ENDPOINT
	}

	return CurrentWeatherService{
		client: c,
		apiURL: apiURL,
		apiKey: apiKey,
		unit:   unit,
		lang:   lang,
	}
}

func (w *CurrentWeatherService) get(url string) (*CurrentWeather, error) {
	response, err := w.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var currentWeather *CurrentWeather
	err = json.NewDecoder(response.Body).Decode(&currentWeather)

	if err != nil {
		return nil, err
	}

	return currentWeather, nil
}

func (w *CurrentWeatherService) CurrentByCityName(cityName string) (*CurrentWeather, error) {
	queryParam := fmt.Sprintf("q=%s&units=%s&appid=%s", url.QueryEscape(cityName), w.unit, w.apiKey)
	requestURL := fmt.Sprintf(w.apiURL, queryParam)

	currentWeather, err := w.get(requestURL)
	if err != nil {
		return nil, err
	}

	return currentWeather, nil
}

func (w *CurrentWeatherService) CurrentByGeoPos(lat float64, lon float64) (*CurrentWeather, error) {
	queryParam := fmt.Sprintf("lat=%v&lon=%v&units=%s&appid=%s", lat, lon, w.unit, w.apiKey)
	requestURL := fmt.Sprintf(w.apiURL, queryParam)

	currentWeather, err := w.get(requestURL)
	if err != nil {
		return nil, err
	}

	return currentWeather, nil
}

func (w *CurrentWeatherService) CurrentByZipCode(zipcode string) (*CurrentWeather, error) {
	queryParam := fmt.Sprintf("zip=%s&units=%s&appid=%s", url.QueryEscape(zipcode), w.unit, w.apiKey)
	requestURL := fmt.Sprintf(w.apiURL, queryParam)

	currentWeather, err := w.get(requestURL)
	if err != nil {
		return nil, err
	}

	return currentWeather, nil
}
