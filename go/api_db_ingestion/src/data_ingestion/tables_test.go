package main

import (
	"fmt"
	"testing"
)

func TestDatabaseUpdateDbTable(t *testing.T) {
	CheckIntegrationTest(t)

	config := GetVariables()
	//needs to be mocked
	var database = PostgreSQL{
		IPAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDB"],
	}

	tables := []string{
		"customers",
		"stores",
	}
	err := UpdateTables(false, tables, &database);if err != nil {
		fmt.Print(err.Error())
	}
}
