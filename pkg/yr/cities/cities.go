package cities

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type City struct {
	CC      string
	City    string
	Country string
	Link    string
}

func ParseCitiesFile() ([]City, error) {
	var cities []City

	reader := csv.NewReader(strings.NewReader(cityList))
	reader.Comma = ';'

	citiesString, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	for _, city := range citiesString {
		if len(city) != 4 {
			return nil, fmt.Errorf("city %v isn't correctly formatted", city)
		}

		cities = append(cities, City{
			CC:      city[0],
			City:    city[1],
			Country: city[2],
			Link:    city[3],
		})
	}

	return cities, nil
}

func CitiesToMap() (map[string][]City, error) {
	cities := make(map[string][]City)

	citiesList, err := ParseCitiesFile()

	if err != nil {
		return nil, err
	}

	for _, city := range citiesList {
		key := strings.ToLower(city.City)
		_, ok := cities[key]

		if ok == true {
			// Multiple cities with the same name
			cities[key] = append(cities[key], city)
			continue
		}

		cities[key] = append([]City{}, city)
	}

	return cities, nil
}
