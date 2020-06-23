package data_ingestion

import (
	"testing"
)

func TestDatabase_UpdateDbTable(t *testing.T) {
	CheckIntegrationTest(t)
	config := get_variables()
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
	update_tables(false, tables, &database )
}
