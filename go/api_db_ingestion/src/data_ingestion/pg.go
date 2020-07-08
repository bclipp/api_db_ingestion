//This module is used for interacting with the postgresql database

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)


type database interface {
	connect() error
	close()
	updateDBTable(table []Row, tableName string) error
	sendQuery(query string)  (sql.Result, error)
	returnTable(tableName string) ([]Row, error)
}

// Database is used to hold the connection related variables
type PostgreSQL struct {
	DB               *sql.DB
	IPAddress        string
	PostgresPassword string
	PostgresUser     string
	PostgresDB       string
}


// Row is used to hold a row of data from a table in the DB
// only common data is used and needed.
type Row struct {
	BlockID   int
	StateCode string
	StateFips int
	BlockPop  int
	ID        int
	Latitude  float64
	Longitude float64
}

// Connect is used to handle connecting to the database
// Params:
// return:
//       error from the connection setup
func (pg PostgreSQL) connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.IPAddress, pg.PostgresUser, pg.PostgresPassword, pg.PostgresDB)

	var err error
	pg.DB, err = sql.Open("postgres", psqlInfo);if err != nil {
		return err
	}

	err = pg.DB.Ping();if err != nil {
		return err
	}

	return nil
}

// Close is used to handle closing the connection to the database
// Params:
// return:
//       error from the connection setup
func (pg PostgreSQL) close() {
	_ = pg.DB.Close()
}

// ReadTable is used for reading data from the database and storing it in the
// table field
// Params:
//       tableName: the table to query
//return:
//       Jason return document
//       rest http response code
//       the error
func (pg PostgreSQL) returnTable(tableName string)([]Row, error) {
	table := make([]Row, 0)
	rows, err := pg.DB.Query(selectTableQuery(tableName, -1));if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			logrus.Printf("error in loadTable - error: %v", err)
		}
	}()

	for rows.Next() {
		var latitude float64

		var longitude float64

		var id int

		err = rows.Scan(&latitude, &longitude);if err != nil {
			return nil, err
		}

		newRow := Row{
			Latitude:  latitude,
			Longitude: longitude,
			ID:        id,
		}
		table = append(table, newRow)
	}

	err = rows.Err();if err != nil {
		return nil, err
	}

	err = rows.Close();if err != nil {
		return nil, err
	}

	return table, nil
}

// SendQuery is used for sending query to a database
// Params:
//       SendQuery: SQL to send
//return:
//		 result variable , see result interface doc in sql
//       the error
func (pg PostgreSQL) sendQuery(query string) (sql.Result, error) {
	result, err := pg.DB.Exec(query)
	if err != nil { return result, err}
	return result, nil
}

// UpdateDbTable is used for taking the the table variable and updating the db
// Params:
//       tableName: the table to query
//return:
//       the error
func (pg PostgreSQL) updateDBTable(table []Row ,tableName string) error {
	for _, row := range table {
		query := updateTableQuery(tableName, row)
		result, err := pg.sendQuery(query);if err != nil {
			return err
		}
		count, err := result.RowsAffected();if err != nil {
			return err
		}

		if count != 1 {
			print("Error when updating row, rows effected is not 1.")
		}
	}

	return nil

}
