package main

import (
	"fmt"
	"testing"
)

func TestPG(t *testing.T) {
	CheckIntegrationTest(t)

	config := GetVariables()

	var pg = PostgreSQL{
		IPAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDB"],
	}

	t.Run("loadTable", func(t *testing.T) {
		err :=pg.connect()
		if err != nil {fmt.Print(err.Error())}
		if err != nil {
			fmt.Print(err.Error())
		}
		defer pg.close()
		var table []Row
		table, err = pg.returnTable("customers")
		if err != nil {fmt.Print(err.Error())}
		if len(table) < 1 {
			t.Errorf("Error, read customers table and no data was found.")
		}
	})
	t.Run("UpdateTable", func(t *testing.T) {
		err := pg.connect()
		if err != nil {fmt.Print(err.Error())}
		defer pg.close()
	})
}