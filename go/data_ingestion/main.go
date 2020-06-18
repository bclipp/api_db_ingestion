package main


func main() {
	config := get_variables()
	print(config)
	tables := [2]string{"customers", "stores"}
	for _, table := range tables {
		UpdateTable(true,table, config)
	}


}

