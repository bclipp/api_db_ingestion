package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetCensusAPI(t *testing.T) {
	// if you need to ignore bad ssl cert
    // http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f
	// add retry for  network issues
	//https://stackoverflow.com/questions/50676817/does-the-http-request-automatically-retry

	lattitude := FloatToString(37.299590)
	longitute := FloatToString(-76.742290)
	_,_,got:= census_api("https://geo.fcc.gov/api/census/area?lat=" + lattitude + "0&lon=" + longitute +"&format=json")
	fmt.Println(" status code: %i", got)

	if ! (got >= 200 && got <= 299) {
		t.Errorf("response code is not in 200 range")
	}
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}