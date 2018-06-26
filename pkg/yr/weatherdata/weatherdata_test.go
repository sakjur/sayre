package weatherdata

import (
	"testing"
	"io/ioutil"
)

func TestWeatherdataParseSample(t *testing.T) {
	bytes, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		println(err)
		t.Fail()
		return
	}

	data := ParseWeatherdata(bytes)

	if data.Location.Name != "Amsterdam" {
		t.Fail()
	}

	for _, time := range data.Forecasts {
		if time.Precipitation.Value != 0.0 {
			t.Fail()
		}
	}
}
