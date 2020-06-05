package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func census_api(url string) (string, error,int){
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, error := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData),error,response.StatusCode
}