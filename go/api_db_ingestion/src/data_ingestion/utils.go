//This module is used for holding reusable functions

package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"testing"
)

// get_variables are used to hold environmental variables read by the app
func GetVariables() map[string]string {
	config := make(map[string]string)
	config["postgresDb"] = os.Getenv("POSTGRES_DB")
	config["postgresUser"] = os.Getenv("POSTGRES_USER")
	config["postgresPassword"] = os.Getenv("POSTGRES_PASSWORD")
	config["IpAddress"] = os.Getenv("DB_IP_ADDRESS")
	config["INT_TEST"] = os.Getenv("INT_TEST")
	return config
}

// FloatToString is used for converting a float to a string
// Params:
//       input_num: used to convert
//return:
//       the string version of the float
func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

// CheckIntergrationTest is used to avoid integration tests if you are running unit tests
func CheckIntegrationTest(t *testing.T) {
	config := GetVariables()
	if config["INT_TEST"] == "" {
		t.Skip("Skipping testing in during unit testing")
	}
}

func setup_log() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	file, err := OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)
}
