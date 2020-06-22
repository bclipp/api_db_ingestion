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
	BlockId string
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
		database.ReadTable(table)
		for _, row := range database.table {
			census, _, err := census_api(row.Latitude, row.Longitude)
			if err != nil {
				return err
			}
			//replace with read table
			table := make([]Row, 0)

			newRow := Row{
				BlockId: census.Results[0].blockId,
				StateCode: census.Results[0].stateCode,
				StateFips: census.Results[0].stateFips,
				BlockPop:  census.Results[0].blockPop,
			}
			table = append(table, newRow)
		}// updating table in db
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

func (d Database) UpdateDbTable(table string, database *Database) error {
	for _, row := range database.table {
		query := update_table_query(table,row)
		result, err := d.Db.Exec(query)
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
