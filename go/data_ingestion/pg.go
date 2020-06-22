package data_ingestion

import (
	"database/sql"
	"fmt"
)

type Database struct {
	Db               *sql.DB
	IpAddress        string
	PostgresPassword string
	PostgresUser     string
	PostgresDb       string
	table            []Row
}

type Row struct {
	BlockId int
	StateCode string
	StateFips int
	BlockPop  int
	Id        int
	Latitude  float64
	Longitude float64
}

func (d Database) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.IpAddress, d.PostgresUser, d.PostgresPassword, d.PostgresDb)
	var err error
	d.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	//defer d.Db.Close()

	err = d.Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (d Database) Close() error {
	return nil
}

func (d Database) ReadTable(tableName string) error {
	table := make([]Row, 0)
	rows, err := d.Db.Query("SELECT *  FROM %s", tableName)
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

func (d Database) SendQuery(query string) (sql.Result, error) {
	result, err := d.Db.Exec(query)
	if err != nil {
		return result, err
	}
	return result,nil
}

func (d Database) UpdateDbTable(table string) error {
	for _, row := range d.table {
		query := update_table_query(table,row)
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
