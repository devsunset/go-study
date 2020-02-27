package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB Example
func main() {
	selectOne()
	selectList()
	insert()
	insertPrepared()
	insertTran()
}

func selectOne() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id string
	err = db.QueryRow("SELECT id FROM temp WHERE id = 1").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func selectList() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var password string
	rows, err := db.Query("SELECT id, password FROM temp where id != ?", 1)
	//MySql ? , Oracle :val1, :val2 , PostgreSQL $1, $2
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, password)
	}
}

func insert() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO temp VALUES (?, ?)", 4, "a")
	if err != nil {
		log.Fatal(err)
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
}

func insertPrepared() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pstmt, err := db.Prepare("UPDATE temp SET id=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer pstmt.Close()

	_, err = pstmt.Exec(5, 4)
	if err != nil {
		log.Fatal(err)
	}
	_, err = pstmt.Exec(6, 5)
	if err != nil {
		log.Fatal(err)
	}
	_, err = pstmt.Exec(7, 6)
	if err != nil {
		log.Fatal(err)
	}
}

func insertTran() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO temp VALUES (?, ?)", 10, "a")
	if err != nil {
		log.Panic(err)
	}

	_, err = tx.Exec("INSERT INTO temp VALUES (?, ?)", 11, "b", "error")
	if err != nil {
		log.Panic(err)
	}

	_, err = tx.Exec("INSERT INTO temp VALUES (?, ?)", 12, "c")
	if err != nil {
		log.Panic(err)
	}
	// 트랜잭션 커밋
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
