package data_ingestion

func update_tabls(concurrent bool, tables []string, database *Database ){


	if concurrent == false {
		for _, table := range tables {
			database.Connect()
			defer database.Db.Close()
			database.ReadTable(table)
			for _, row := range database.table {
				println(row)

				response,_,_ := census_api(row.Latitude, row.Longitude)
				row.BlockId = response.Results[0].blockId
				row.BlockPop = response.Results[0].blockPop
				row.StateCode = response.Results[0].stateCode
				row.StateFips = response.Results[0].stateFips
			}
			database.UpdateTable(table,database)
		}

	}

}
