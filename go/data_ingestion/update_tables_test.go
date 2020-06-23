package data_ingestion

import (
	"fmt"
	"testing"
)

func TestDatabaseUpdateDbTable(t *testing.T) {
	CheckIntegrationTest(t)
	config := getVariables()
	//needs to be mocked
	var database = Database{
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
	if err != nil {
		fmt.Print(err.Error())
	}
}
