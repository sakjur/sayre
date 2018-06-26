package weatherdata

import (
	"encoding/xml"
)

type Weatherdata struct {
	XMLName   xml.Name   `xml:"weatherdata"`
	Location  Location   `xml:"location"`
	Sun       Sun        `xml:"sun"`
	Forecasts []Forecast `xml:"forecast>tabular>time"`
	Credit    CreditLink `xml:"credit>link"`
}

type Location struct {
	Name    string `xml:"name"`
	Type    string `xml:"type"`
	Country string `xml:"country"`
}

type Sun struct {
	Rise string `xml:"rise,attr"`
	Set  string `xml:"set,attr"`
}

type Forecast struct {
	From          string        `xml:"from,attr"`
	To            string        `xml:"to,attr"`
	Precipitation Precipitation `xml:"precipitation"`
	Temperature   Temperature   `xml:"temperature"`
	Symbol        SymbolXml     `xml:"symbol"`
	WindDirection WindDirection `xml:"windDirection"`
	WindSpeed     WindSpeed     `xml:"windSpeed"`
}

type SymbolXml struct {
	Number int    `xml:"numberEx,attr"`
	Name   string `xml:"name,attr"`
}

type WindDirection struct {
	Deg  float64 `xml:"deg,attr"`
	Code string  `xml:"code,attr"`
	Name string  `xml:"name,attr"`
}

type WindSpeed struct {
	MPS  float64 `xml:"mps,attr"`
	Name string  `xml:"name,attr"`
}

type Precipitation struct {
	Value float64 `xml:"value,attr"`
}

type Temperature struct {
	Unit  string  `xml:"unit,attr"`
	Value float64 `xml:"value,attr"`
}

type CreditLink struct {
	Text string `xml:"text,attr"`
	Url  string `xml:"url,attr"`
}

func ParseWeatherdata(bytes []byte) Weatherdata {
	var data Weatherdata

	err := xml.Unmarshal(bytes, &data)
	if err != nil {
		println(err)
	}

	return data
}
