package main

import (
	"fmt"
	"testing"
)

func TestGetCensusAPI(t *testing.T) {

	lattitude := FloatToString(37.299590)
	longitute := FloatToString(-76.742290)
	_,_,got:= census_api("https://geo.fcc.gov/api/census/area?lat=" + lattitude + "0&lon=" + longitute +"&format=json")
	fmt.Println(" status code: %i", got)

	if ! (got >= 200 && got <= 299) {
		t.Errorf("response code is not in 200 range")
	}
}