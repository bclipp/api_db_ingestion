//This module is used for controling the lookup and update process

package data_ingestion

// update_tables is used for handling the update process
// table field
// Params:
//       concurrent: used to control if the process is done concurrently
//       tables: what tables are going to be updated
// 		 database: the struct used for handling databses.
//return:
//       the error
func update_tables(concurrent bool, tables []string, database *Database) error {
	if !concurrent {
		for _, table := range tables {
			err := database.Connect()
			if err != nil {
				return err
			}
			defer database.Db.Close()
			err = database.ReadTable(table)
			if err != nil {
				return err
			}
			for _, row := range database.table {
				response, _, err := census_api(row.Latitude, row.Longitude)
				if err != nil {
					return err
				}
				row.BlockId = response.Results[0].BlockId
				row.BlockPop = response.Results[0].BlockPop
				row.StateCode = response.Results[0].StateCode
				row.StateFips = response.Results[0].StateFips
			}
			err = database.UpdateDbTable(table)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
