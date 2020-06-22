package data_ingestion

func update_tables(concurrent bool, tables []string, database *Database )error{
	if concurrent == false {
		for _, table := range tables {
			database.Connect()
			defer database.Db.Close()
			database.ReadTable(table)
			for _, row := range database.table {
				response,_,err := census_api(row.Latitude, row.Longitude)
				if err != nil {
					return err
				}
				row.BlockId = response.Results[0].blockId
				row.BlockPop = response.Results[0].blockPop
				row.StateCode = response.Results[0].stateCode
				row.StateFips = response.Results[0].stateFips
			}
			database.UpdateDbTable(table)
		}
	}
	return nil
}
