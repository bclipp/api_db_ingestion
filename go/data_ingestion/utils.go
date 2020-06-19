package data_ingestion

import (
	"os"
	"strconv"
	"testing"
)

func get_variables()  map[string]string{
	config := make(map[string]string)
	config["postgresDb"] = os.Getenv("POSTGRES_DB")
	config["postgresUser"] = os.Getenv("POSTGRES_USER")
	config["postgresPassword"] = os.Getenv("POSTGRES_PASSWORD")
	config["IpAddress"] = os.Getenv("DB_IP_ADDRESS")
	config["INT_TEST"] = os.Getenv("INT_TEST")
	return config
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func CheckIntergrationTest(t *testing.T){
	config := get_variables()
	if config["INT_TEST"] == "" {
		t.Skip("Skipping testing in during unit testing")
	}
}