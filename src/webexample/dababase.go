package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:admin@(127.0.0.1)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Drop Table
	{
		_, err := db.Exec(`DROP TABLE members`)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create a new table
	{
		query := `
			CREATE TABLE members(
				id INT AUTO_INCREMENT,
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY (id)		
			);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	// Insert a new user
	{
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO members(username,password,created_at) VALUES(?,?,?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	// Query a single user
	{
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username,password,created_at FROM members WHERE  id= ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, password, createdAt)
	}

	// Query all members
	{
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}
		rows, err := db.Query(`SELECT id, username,password,created_at FROM members`)
		if err != nil {
			log.Fatal(err)
		}

		var members []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}

			members = append(members, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", members)
		fmt.Println()
	}

	// Delete a row
	{
		_, err := db.Exec(`DELETE FROM members WHERE id= ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
