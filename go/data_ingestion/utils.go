package main

import (
	"os"
	"strconv"
)

func get_variables()  map[string]string{
	config := make(map[string]string)
	config["postgresDb"] = os.Getenv("POSTGRES_DB")
	config["postgresUser"] = os.Getenv("POSTGRES_USER")
	config["postgresPassword"] = os.Getenv("POSTGRES_PASSWORD")
	config["IpAddress"] = os.Getenv("DB_IP_ADDRESS")
	return config
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}