package data_ingestion

import (
	"testing"
)

func TestPG(t *testing.T) {
	CheckIntergrationTest(t)
	config := get_variables()
	var database = Database{
		IpAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDb:       config["postgresDb"],
	}

	t.Run("Read", func(t *testing.T) {
		database.Connect()
		defer database.Db.Close()
		database.ReadTable("customers")
		if len(database.table) < 1 {
			t.Errorf("Error, read customers table and no data was found.")
		}
	})
	t.Run("UpdateTable", func(t *testing.T) {
		database.Connect()
		defer database.Db.Close()
	})

}