package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "log"

func main() {
	db, err := sql.Open("mysql", "root:@/gleaner")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}
