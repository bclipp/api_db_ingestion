package main

import (
	"github.com/bclipp/api_db_ingestion"
)

func main() {
	config := data_ingestion.get_variables()
	print(config)
	tables := [2]string{"customers", "stores"}
	for _, table := range tables {
		data_ingestion.UpdateTable(true,table, config)
	}
}

