package main

import (
	"fmt"
	"testing"
)

func TestGetCensusAPIStatus(t *testing.T) {
	// if you need to ignore bad ssl cert
    // http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f
	// add retry for  network issues
	//https://stackoverflow.com/questions/50676817/does-the-http-request-automatically-retry

	lattitude := 37.299590
	longitute := -76.742290
	_,got,_:= census_api(lattitude,longitute)
	fmt.Println(" status code: %i", got)
	if ! (got >= 200 && got <= 299) {
		t.Errorf("response code is not in 200 range")
	}
}

func TestGetCensusAPIBody(t *testing.T) {

	lattitude := 37.299590
	longitute := -76.742290
	census,_,_:= census_api(lattitude,longitute)
	if ! (len(census.Results) > 0) {
		t.Errorf("block id is not present")
	}
}
