package main

import (
	"database/sql"
	"fmt"
)

type Database struct {
	Db *sql.DB
	IpAddress string
	PostgresPassword string
	PostgresUser string
	PostgresDb string
	table []Row
}

type Row struct {
	BlockFips string
	StateCode string
	StateFips string
	BlockPop string
	Id int
	Latitude float64
	Longitude float64
}

func UpdateTable(serial bool,config map[string]string) error {
	if serial == true {
		//connect
		//read table
		//interate on table updating each row
		//disconnect
	} else {
		//add concurency
	}

	return nil
}

func (d Database) Connect(config map[string]string) error {
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config["IpAddress"], config["postgresUser"], config["postgresPassword"], config["postgresDb"])
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


func (d Database) Read(tableName string)error{

	table :=  make([]Row, 0)
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
		census, _, _ := census_api(latitude,longitude)
		newRow := Row{
			Latitude : latitude ,
			Longitude : longitude ,
			Id : id ,
			BlockFips : census.Results[0].blockFips,
			StateCode : census.Results[0].blockFips,
			StateFips : census.Results[0].blockPop,
			BlockPop : census.Results[0].blockPop,
		}
		table = append(table, newRow)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func (d Database) Update() error{
	return nil
}
