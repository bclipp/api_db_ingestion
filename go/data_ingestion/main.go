package main

import (
	"fmt"
	"strconv"
)


func main() {
	lattitude := FloatToString(37.299590)
	longitute := FloatToString(-76.742290)
	reponse,_,_ := census_api("https://geo.fcc.gov/api/census/area?lat=" + lattitude + "0&lon=" + longitute +"&format=json")
	fmt.Println(reponse)

}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}