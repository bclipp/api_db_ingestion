package data_ingestion

import "fmt"

func update_table_query(table string, row Row)string{
	state_fips := row.StateFips
	state_code := row.StateCode
	block_pop := row.BlockPop
	block_id := row.BlockId
	table_id := row.Id
	return fmt.Sprintf(
		"UPDATE %s SET state_fips = %d, state_code = '%s', block_pop = %d, block_id = %d WHERE ID = %d;",
		table,state_fips,state_code,block_pop,block_id,table_id)
}