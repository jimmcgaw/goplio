package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "log"

func main() {
	db, err := sql.Open("mysql", "root:@/gleaner")
	if err != nil {
		fmt.Println("Error opening connection to MySQL")
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(db)

	// fetch all user records and display their data
	rows, err := db.Query("SELECT id, email, username FROM auth_user;")
	if err != nil {
		fmt.Println("Error fetching user data")
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)

	var (
		id       int
		username string
		email    string
	)
	for rows.Next() {
		err := rows.Scan(&id, &email, &username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, email, username)
	}

	// we've got the last user ID, let's use it to make a single row query
	log.Println(id)

	var (
		firstName string
		lastName  string
	)

	err = db.QueryRow("SELECT first_name, last_name FROM auth_user where id = ?", id).Scan(&firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User's full name is", firstName, lastName)

}
