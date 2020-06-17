package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


var httpClient = &http.Client{Timeout: 10 * time.Second}

type Inner struct {
	blockFips string `json:"block_fips"`
	stateCode string `json:"state_code"`
	stateFips string `json:"state_fips"`
	blockPop string `json:"block_pop_2015"`
}
type Outer struct {
	Results []Inner `json:"results"`
}



func census_api( latitude,longitude float64) (Outer,int,error){
	url := "https://geo.fcc.gov/api/census/area?lat=" + FloatToString(latitude) + "0&lon=" + FloatToString(longitude) + "&format=json"
	response, error := httpClient.Get(url)
	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}

	var census Outer
	json.Unmarshal(body, &census)
	fmt.Printf("%+v\n", census)
	return census,response.StatusCode,error
}

