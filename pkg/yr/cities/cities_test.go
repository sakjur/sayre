package cities

import (
	"log"
	"testing"
)

func TestParseCitiesFile(t *testing.T) {
	_, err := ParseCitiesFile()

	if err != nil {
		t.Fail()
	}
}

func TestCitiesToMap(t *testing.T) {
	cityMap, err := CitiesToMap()

	if err != nil {
		log.Println(err)
		t.Fail()
	}

	city := cityMap["stockholm"]
	if city[0].City != "Stockholm" {
		t.Fail()
	}
	if city[0].Country != "Sweden" {
		t.Fail()
	}
	if city[0].CC != "SE" {
		t.Fail()
	}
}

func TestCitiesToMapWithMultipleResults(t *testing.T) {
	cityMap, err := CitiesToMap()

	if err != nil {
		log.Println(err)
		t.Fail()
	}

	cities := cityMap["springfield"]

	if len(cities) != 3 {
		t.Fail()
	}

	for _, city := range cities {
		if city.City != "Springfield" {
			t.Fail()
		}
	}
}
