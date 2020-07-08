//This module is used for holding reusable functions

package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"testing"
	"time"
)

// get_variables are used to hold environmental variables read by the app
func GetVariables() map[string]string {
	config := make(map[string]string)
	config["postgresDb"] = os.Getenv("POSTGRES_DB")
	config["postgresUser"] = os.Getenv("POSTGRES_USER")
	config["postgresPassword"] = os.Getenv("POSTGRES_PASSWORD")
	config["postgresIP"] = os.Getenv("POSTGRES_IP_ADDRESS")
	config["intergrationTest"] = os.Getenv("INT_TEST")

	return config
}

// FloatToString is used for converting a float to a string
// Params:
//       input_num: used to convert
//return:
//       the string version of the float
func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

// CheckIntergrationTest is used to avoid integration tests if you are running unit tests
func CheckIntegrationTest(t *testing.T) {
	config := GetVariables()
	if config["intergrationTest"] != "True" {
		t.Skip("Skipping testing in during unit testing")
	}
}

func SetupLog()error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)

	logFileName := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); if err != nil {
		return err
	}

	logrus.SetOutput(file)

	return nil
}
