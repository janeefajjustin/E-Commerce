package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Initialize() *sql.DB{
	err := OpenDatabase()
	fmt.Println("DB Initializing....")
	if err != nil {
		log.Printf("A new error %v", err)
	}
	//defer CloseDatabase()
	fmt.Println("Done...")
	return DB
}

func CreateTable() {
	createUserTable :=
		`CREATE TABLE IF NOT EXISTS user(
     userid serial,
	 username varchar(100) unique,
	 password varchar(100),
	 phonenumber int64 unique,
	 email varchar(100) unique,
	 firstname varchar(100),
	 middlename varchar(100),
	 lastname varchar(100)
	);`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		log.Printf("A new error %v", err)
	}
}


func OpenDatabase() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=8976 dbname=Ecommerce sslmode=disable")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
	return nil

}

func CloseDatabase() error {
	return DB.Close()
}

