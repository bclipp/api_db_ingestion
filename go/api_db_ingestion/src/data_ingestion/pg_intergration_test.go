package main

import (
	"database/sql"
	"fmt"
	"testing"
)


func Test1(t *testing.T) {
	config := GetVariables()

	var pg = PostgreSQL{
		IPAddress:        config["postgresIP"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDB:       config["postgresDB"],
	}

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		pg.PostgresUser,
		pg.PostgresPassword,
		pg.IPAddress,
		pg.PostgresDB)

	db, err := sql.Open("postgres", psqlInfo);if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM customers;")
	defer rows.Close()

	/*DB, err := pg.connect();if err != nil {
		fmt.Print(err.Error())
	}
	DB.Stats()

	defer pg.close()
*/
	//got,_ := pg.sendQuery("SELECT * FROM CUSTOMERS;")
	//println(got)

}

/*
func TestPG(t *testing.T) {
	//CheckIntegrationTest(t)
	t.Run("loadTable", func(t *testing.T) {
		config := GetVariables()
		var pg = PostgreSQL{
			IPAddress:        config["postgresIP"],
			PostgresPassword: config["postgresPassword"],
			PostgresUser:     config["postgresUser"],
			PostgresDB:       config["postgresDB"],
		}

		err := pg.connect();if err != nil {
			fmt.Print(err.Error())
		}
		defer pg.close()
		var table []Row
		table, err = pg.returnTable("customers", 1)
		if err != nil {fmt.Print(err.Error())}
		if len(table) < 1 {
			t.Errorf("Error, read customers table and no data was found.")
		}
	})
	t.Run("UpdateTable", func(t *testing.T) {
		config := GetVariables()
		var pg = PostgreSQL{
			IPAddress:        config["IpAddress"],
			PostgresPassword: config["postgresPassword"],
			PostgresUser:     config["postgresUser"],
			PostgresDB:       config["postgresDB"],
		}
		err := pg.connect()
		if err != nil {fmt.Print(err.Error())}
		defer pg.close()
	})
}
*/