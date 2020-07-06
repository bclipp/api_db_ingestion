package main

import (
	"fmt"
	"testing"
)

func TestGetCensusAPIStatus(t *testing.T) {
	CheckIntegrationTest(t)

	latitude := 37.299590
	longitude := -76.742290
	_, got, _ := censusAPI(latitude, longitude)
	fmt.Println(" status code: %i", got)

	if !(got >= 200 && got <= 299) {
		t.Errorf("response code is not in 200 range")
	}
}

func TestGetCensusAPIBody(t *testing.T) {
	CheckIntegrationTest(t)

	latitude := 37.299590
	longitute := -76.742290
	census, _, _ := censusAPI(latitude, longitute)

	if !(len(census.Results) > 0) {
		t.Errorf("block id is not present")
	}
}
