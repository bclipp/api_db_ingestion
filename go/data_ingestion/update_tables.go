package data_ingestion

func update_tabls(concurrent bool, tables []string, database *Database ){


	if concurrent == false {
		for _, table := range tables {
			database.Connect()
			defer database.Db.Close()
			database.ReadTable(table)
			for _, row := range database.table {
				println(row)

				// lookup each row
				// updating table in db
			}
		}

	}

}
