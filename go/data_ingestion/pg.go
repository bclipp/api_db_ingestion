//This module is used for interacting with the postgresql database

package data_ingestion

import (
	"database/sql"
	"fmt"
)

// Database is used to hold the connection related variables
type Database struct {
	Db               *sql.DB
	IpAddress        string
	PostgresPassword string
	PostgresUser     string
	PostgresDb       string
	table            []Row
}

// Row is used to hold a row of data from a table in the DB
// only common data is used and needed.
type Row struct {
	BlockId   int
	StateCode string
	StateFips int
	BlockPop  int
	Id        int
	Latitude  float64
	Longitude float64
}

// Connect is used to handle connecting to the database
// Params:
// return:
//       error from the connection setup
func (d Database) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.IpAddress, d.PostgresUser, d.PostgresPassword, d.PostgresDb)
	var err error
	d.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = d.Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Close is used to handle closing the connection to the database
// Params:
// return:
//       error from the connection setup
func (d Database) Close() error {
	return nil
}

// ReadTable is used for reading data from the database and storing it in the
// table field
// Params:
//       tableName: the table to query
//return:
//       Jason return document
//       rest http response code
//       the error
func (d Database) ReadTable(tableName string) error {
	table := make([]Row, 0)
	rows, err := d.Db.Query(select_table(tableName, -1))
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var latitude float64
		var longitude float64
		var id int
		err = rows.Scan(&latitude, &longitude)
		if err != nil {
			return err
		}
		newRow := Row{
			Latitude:  latitude,
			Longitude: longitude,
			Id:        id,
		}
		table = append(table, newRow)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	d.table = table
	return nil
}

// SendQuery is used for sending query to a database
// Params:
//       SendQuery: SQL to send
//return:
//		 result variable , see result interface doc in sql
//       the error
func (d Database) SendQuery(query string) (sql.Result, error) {
	result, err := d.Db.Exec(query)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateDbTable is used for taking the the table variable and updating the db
// Params:
//       tableName: the table to query
//return:
//       the error
func (d Database) UpdateDbTable(tableName string) error {
	for _, row := range d.table {
		query := update_table_query(tableName, row)
		result, err := d.SendQuery(query)
		if err != nil {
			return err
		}
		count, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if count != 1 {
			print("Error when updating row, rows effected is not 1.")
		}
	}
	return nil
}
