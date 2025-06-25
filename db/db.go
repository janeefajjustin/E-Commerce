package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initialize() (*sql.DB, error) {
	err := OpenDatabase()
	fmt.Println("DB Initializing....")
	if err != nil {
		return nil, err
	}
	//defer CloseDatabase()
	fmt.Println("DB Initialized successfully")
	return DB, nil
}

func CreateTable() {
	createUserTable :=
		`CREATE TABLE IF NOT EXISTS users(
     userid serial,
	 username varchar(100) unique,
	 password varchar(100),
	 phonenumber varchar(20) unique,
	 email varchar(100) unique,
	 firstname varchar(100),
	 middlename varchar(100),
	 lastname varchar(100)
	)`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		log.Printf("A new error %v", err)
	}

	createProductCategoryTable :=
		`CREATE TABLE IF NOT EXISTS product_category(
	categoryid serial,
	categoryname varchar(100),
    createdat timestamp default(null),
	updatedat timestamp default(null),
	deletedat timestamp default(null)
	)`

	_, err = DB.Exec(createProductCategoryTable)
	if err != nil {
		log.Printf("A new error %v", err)
	   
	}

	createProductImageTable:=
	`CREATE TABLE IF NOT EXISTS product_image(

    image_id serial,

    product_id integer,
    image_url varchar(200),

    createdat timestamp default(null),

    updatedat timestamp default(null),

    deletedat timestamp default(null)

    )`
	_, err = DB.Exec(createProductImageTable)
	if err != nil {
		log.Printf("A new error %v", err)
	   
	}

	createProductSizeTable:= `CREATE TABLE IF NOT EXISTS product_size(
size_id serial,
size varchar(10),
product_id integer,
product_quantity integer,
createdat timestamp default(null),
updatedat timestamp default(null),
deletedat timestamp default(null)
)`
_, err = DB.Exec(createProductSizeTable)
	if err != nil {
		log.Printf("A new error %v", err)
	   
	}

	createProductTable:=`CREATE TABLE IF NOT EXISTS product(
product_id serial,
  product_name varchar(10),
  createdat timestamp default(null),
  updatedat timestamp default(null),
  deletedat timestamp default(null)
)`
_, err = DB.Exec(createProductTable)
	if err != nil {
		log.Printf("A new error %v", err)
	   
	}
}

func OpenDatabase() error {
	var err error
	// connectionstring:=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=desable",)
	DB, err = sql.Open("postgres", "user=postgres password=8976 dbname=Ecommerce sslmode=disable")
	if err != nil {
		return err
	}

	// if err != nil {
	// 	return err
	// }

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
	return nil

}

func CloseDatabase() error {
	return DB.Close()
}
