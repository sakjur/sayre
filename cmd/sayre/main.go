package main

import (
	"github.com/sakjur/sayre/pkg/yr"
	"flag"
	"github.com/sakjur/sayre/pkg/yr/cities"
	"strings"
)

func main() {
	var city string
	flag.StringVar(&city, "city", "amsterdam", "City to access weather data from")
	flag.Parse()

	city = strings.ToLower(city)

	citiesMap, err := cities.CitiesToMap()
	if err != nil {
		println(err)
		return
	}

	weatherdata := citiesMap[city]

	for _, c := range weatherdata {
		data, _ := yr.Fetch(c.Link)

		println(data.ToSlack())
	}

}
