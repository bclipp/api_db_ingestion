package main

import (
	"github.com/bclipp/api_db_ingestion/go/data_ingestion"
	"fmt"
)

func main() {

	config := data_ingestion.GetVariables()
	//needs to be mocked
	var database = data_ingestion.Database{
		IpAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDb:       config["postgresDb"],
	}
	tables := []string{
		"customers",
		"stores",
	}
	err := data_ingestion.UpdateTables(false, tables, &database)
	if err != nil {fmt.Print(err.Error())}

}