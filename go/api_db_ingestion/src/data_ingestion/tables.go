//This module is used for controlling the lookup and update process

package main

import (
	"github.com/sirupsen/logrus"
)

// update_tables is used for handling the update process
// table field
// Params:
//       concurrent: used to control if the process is done concurrently
//       tables: what tables are going to be updated
// 		 database: the struct used for handling databses.
//return:
//       the error
func UpdateTables(concurrent bool, tables []string, db database) error {
	if !concurrent {
		for _, tableName := range tables {
			contextLogger := logrus.WithFields(logrus.Fields{
				"tableName": tableName,
			})
			contextLogger.Debug("Starting Data Import Loop")

			_, err := db.connect();if err != nil {
				logrus.Println(err)
				return err
			}

			defer db.close()

			contextLogger.Debug("Loading the Table from DataBase")

			table ,err := db.returnTable(tableName, -1)
			if err != nil {
				logrus.Println(err)
				return err
			}

			contextLogger.Debug("Looking up Census Data")

			for _, row := range table {
				response, _, err := censusAPI(row.Latitude, row.Longitude);if err != nil {
					return err
				}

				row.BlockID = response.Results[0].BlockID
				row.BlockPop = response.Results[0].BlockPop
				row.StateCode = response.Results[0].StateCode
				row.StateFips = response.Results[0].StateFips
			}

			err = db.updateDBTable(table, tableName);if err != nil {
				return err
			}
		}
	}

	return nil
}
