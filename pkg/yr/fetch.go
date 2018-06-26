package yr

import (
	"github.com/sakjur/sayre/pkg/yr/weatherdata"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func FetchHourByHour(baseCity string) (weatherdata.Weatherdata, error) {
	url := strings.Replace(baseCity, "forecast.xml", "forecast_hour_by_hour.xml", 1)
	return fetch_weatherdata(url)
}

func Fetch(baseCity string) (weatherdata.Weatherdata, error) {
	return fetch_weatherdata(baseCity)
}

func fetch_weatherdata(url string) (weatherdata.Weatherdata, error) {
	client := http.Client{
		Timeout: 20 * time.Second,
	}

	result, err := client.Get(url)

	if err != nil {
		println(err)
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)

	if err != nil {
		println(err)
	}

	data := weatherdata.ParseWeatherdata(body)

	return data, nil
}
