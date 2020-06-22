package data_ingestion

import "fmt"

func update_table_query(table string, row Row)string{
	state_fips := row.StateFips
	state_code := row.StateCode
	block_pop := row.BlockPop
	block_id := row.BlockId
	table_id := row.Id
	return fmt.Sprintln(`UPDATE
	%v
	SET
	state_fips = %v, state_code = '%v', block_pop = %v, block_id = %v
	WHERE
	ID = %v;`,table,state_fips,state_code,block_pop,block_id,table_id)
}