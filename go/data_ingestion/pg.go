package main

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
	BlockFips string
	StateCode string
	StateFips string
	BlockPop  string
	Id        int
	Latitude  float64
	Longitude float64
}

func UpdateTable(serial bool, table string, config map[string]string) error {
	if serial == true {
		var database = Database{
			IpAddress:        config["IpAddress"],
			PostgresPassword: config["postgresPassword"],
			PostgresUser:     config["postgresUser"],
			PostgresDb:       config["postgresDb"],
		}
		database.Connect()
		defer database.Db.Close()
		database.Read(table)
		for _, row := range database.table {
			census, _, err := census_api(row.Latitude, row.Longitude)
			if err != nil {
				return err
			}
			table := make([]Row, 0)
			newRow := Row{
				BlockFips: census.Results[0].blockFips,
				StateCode: census.Results[0].blockFips,
				StateFips: census.Results[0].blockPop,
				BlockPop:  census.Results[0].blockPop,
			}
			table = append(table, newRow)

		}
	} else {
		//add concurency
	}

	return nil
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

func (d Database) Read(tableName string) error {

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

func (d Database) Update(table string, row Row) error {
	sqlQuery := `
UPDATE $2
SET BlockFips = $3, StateCode = $4, StateFips = $5, BlockPop = $6
WHERE id = $1;`
	result, err := d.Db.Exec(
		sqlQuery,
		row.Id, table,
		row.BlockFips,
		row.StateCode,
		row.StateFips,
		row.BlockPop)
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
	return nil
}
