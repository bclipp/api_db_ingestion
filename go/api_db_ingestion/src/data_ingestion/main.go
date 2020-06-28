package main

import (
	"fmt"
)

func main() {

	config := GetVariables()
	//needs to be mocked
	var database = Postgresql{
		IpAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDb:       config["postgresDb"],
	}
	tables := []string{
		"customers",
		"stores",
	}
	err := UpdateTables(false, tables, &database)
	if err != nil {fmt.Print(err.Error())}

}