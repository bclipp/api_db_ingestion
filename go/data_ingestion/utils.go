package main

import "os"

func get_variables()  [4]string{
	postgresDb := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	IpAddress :=os.Getenv("DB_IP_ADDRESS")
	config := [4]string{postgresDb, postgresUser, postgresPassword, IpAddress}
	return config
}
