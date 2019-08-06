package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParseCityProfile(t *testing.T) {
	file, e := ioutil.ReadFile("profileCityTest.html")
	if e != nil {
		log.Println(e.Error())
	}

	parseResult := ParseCityProfile(file)

	for _, item := range parseResult.Items {
		log.Printf("result: %v", item)

	}

}
