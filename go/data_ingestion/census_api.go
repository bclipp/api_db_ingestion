package data_ingestion

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
	blockId int `json:"block_id"`
	stateCode string `json:"state_code"`
	stateFips int `json:"state_fips"`
	blockPop int `json:"block_pop_2015"`
}
type Outer struct {
	Results []Inner `json:"results"`
}



func census_api( latitude,longitude float64) (Outer,int,error){
	url := "https://geo.fcc.gov/api/census/area?lat=" + FloatToString(latitude) + "0&lon=" + FloatToString(longitude) + "&format=json"
	response, err := httpClient.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var census Outer
	json.Unmarshal(body, &census)
	fmt.Printf("%+v\n", census)
	return census,response.StatusCode,err
}

