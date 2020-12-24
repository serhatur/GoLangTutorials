package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// dbDriver := "mysql"
	// dbUser := "u9115920_user122"
	// dbPass := "NQax91N3HZni45O"
	// dbName := "u9115920_db122"
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(94.73.147.224)/"+dbName)
	db, err := sql.Open("mysql", "merym_golang:NQax91N3HZni45O@tcp(94.73.147.224)/u9115920_db122")
	if err != nil {
		panic(err.Error)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error)
	}

	//Veri ekleme
	res, err := db.Exec("INSERT INTO users (`Username`, `Email`) VALUES ('sibel', 'sibel@hotmail.com')")
	if err != nil {
		panic(err.Error)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		panic(err.Error)
	}
	log.Printf("Inserted %d rows", rowCount)

	var (
		ID       int
		Username string
		Email    string
	)

	//Veri okuma
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error)
	}

	for rows.Next() {
		err = rows.Scan(&ID, &Username, &Email)
		if err != nil {
			panic(err.Error)
		}
		log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email)
	}
	//rows.Close()

	rows2, err := db.Query("SELECT * FROM users WHERE ID = ?", 1)
	if err = rows2.Err(); err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()
	for rows2.Next() {
		err = rows2.Scan(&ID, &Username, &Email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bulunan satır içeriği ID : %q", strconv.Itoa(ID)+" "+Username+" "+Email)
	}
	err = rows2.Err()
	if err != nil {
		log.Fatal(err)
	}

	errErr := db.QueryRow("select * from users limit 1").Scan(&ID, &Username, &Email)
	if errErr != nil {
		if errErr == sql.ErrNoRows {

		} else {
			log.Fatal(errErr)
		}
	}
	log.Println("Bir satır bulundu : ", ID)

	//Multiple active result set : MARS
	//_, err := db.Exec("DELETE FROM users","DELETE FROM users")

	//Preparing queries
	stmt, errQ := db.Prepare("select * from users where ID = ?")
	if errQ != nil {
		log.Fatal(errQ)
	}
	defer stmt.Close()
	rows, errX := stmt.Query(7)
	if errX != nil {
		log.Fatal(errX)
	}
	defer rows.Close()
	for rows.Next() {
		prScan := rows.Scan(&ID, &Username, &Email)
		if prScan != nil {
			log.Fatal(prScan)
		}
		log.Printf("Bulunan satır içeriği ID : %q", strconv.Itoa(ID)+" "+Username+" "+Email)
	}
}
