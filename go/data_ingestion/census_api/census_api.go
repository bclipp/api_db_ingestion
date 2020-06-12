package census_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


var httpClient = &http.Client{Timeout: 10 * time.Second}





func census_api(url string) error{
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

	type Inner struct {
		Key2 string `json:"block_fips"`
		Key3 string `json:"block_pop_2015"`
		Key4 string `json:"state_code"`
	}
	type Outer struct {
		Key1 []Inner `json:"results"`
	}
	var cont Outer
	json.Unmarshal(body, &cont)
	fmt.Printf("%+v\n", cont)
	return nil
}

