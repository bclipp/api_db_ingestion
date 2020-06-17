package main

import (
	"database/sql"
	"fmt"
)

type databaseEnv struct {
	db *sql.DB
	age  int
}

type Row struct {
	blockFips string
	stateCode string
	stateFips string
	blockPop string
	id int
	latitude float64
	longitude float64
}

func read(config map[string]string, tableName string)[]Row{

	table :=  make([]Row, 0)


	database := new(databaseEnv)
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config["IpAddress"], config["postgresUser"], config["postgresPassword"], config["postgresDb"])

	database.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer database.db.Close()

	err = database.db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := database.db.Query("SELECT *  FROM %s", tableName)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var latitude float64
		var longitude float64
		var id int
		err = rows.Scan(&latitude, &longitude)
		if err != nil {
			// handle this error
			panic(err)
		}
		census, _, _ := census_api(latitude,longitude)
		newRow := Row{
			latitude : latitude ,
			longitude : longitude ,
			id : id ,
			blockFips : census.Results[0].blockFips,
			stateCode : census.Results[0].blockFips,
			stateFips : census.Results[0].blockPop,
			blockPop : census.Results[0].blockPop,
		}
		table = append(table, newRow)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return table
}

//update

//lookup_row
