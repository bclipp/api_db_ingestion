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
	BlockFips string `json:"block_fips"`
	StateCode string `json:"state_code"`
	StateFips string `json:"state_fips"`
	BlockPop string `json:"block_pop_2015"`
}
type Outer struct {
	Results []Inner `json:"results"`
}



func census_api(url string) (Outer,int,error){
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

