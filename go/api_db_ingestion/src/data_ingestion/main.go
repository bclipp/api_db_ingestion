package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init(){
	setup_log()
}

func main() {

	config := GetVariables()
	//needs to be mocked
	var pg = PostgreSQL{
		IpAddress:        config["IpAddress"],
		PostgresPassword: config["postgresPassword"],
		PostgresUser:     config["postgresUser"],
		PostgresDb:       config["postgresDb"],
	}
	tables := []string{
		"customers",
		"stores",
	}
	err := UpdateTables(false, tables, &pg)
	if err != nil {log.Fatal(err.Error())}

}