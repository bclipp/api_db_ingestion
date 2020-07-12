//This module is used for controlling the lookup and update process

package main

import (
	EasyDatabase "github.com/bclipp/EasyDatabase"
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
func UpdateTables(concurrent bool, tables []string, database EasyDatabase.PostgreSQL) error {
	if !concurrent {
		for _, tableName := range tables {
			contextLogger := logrus.WithFields(logrus.Fields{
				"tableName": tableName,
			})
			contextLogger.Debug("Starting Data Import Loop")

			err := database.Connect();if err != nil {
				logrus.Println(err)
				return err
			}

			defer database.Close()

			contextLogger.Debug("Loading the Table from DataBase")

			table ,err := database.ReturnTable(tableName, -1)
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

			err = database.UpdateDBTable(table, tableName);if err != nil {
				return err
			}
		}
	}

	return nil
}
