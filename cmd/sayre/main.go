package main

import (
	"github.com/sakjur/sayre/pkg/yr"
	"flag"
	"github.com/sakjur/sayre/pkg/yr/cities"
)

func main() {
	city := flag.String("city", "amsterdam", "City to access weather data from")

	citiesMap, err := cities.CitiesToMap()
	if err != nil {
		println(err)
		return
	}

	weatherdata := citiesMap[*city]

	for _, c := range weatherdata {
		data, _ := yr.Fetch(c.Link)

		println(data.ToSlack())
	}

}
