package main

import (
	EasyDatabase "github.com/bclipp/EasyDatabase"
	"github.com/sirupsen/logrus"
)


func init() {
	SetupLog()
}
func main() {
	config := GetVariables()
	//needs to be mocked
	var pg = EasyDatabase.PostgreSQL{
		IPAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDb"],
	}

	tables := []string{
		"customers",
		"stores",
	}
	 err := UpdateTables(false, tables, &pg)

	if err != nil {
		logrus.Fatal(err.Error())
	}
}
