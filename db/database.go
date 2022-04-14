package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3307)/databaseName

const url = "root:1234@tcp(localhost:3307)/gallery"

var db *sql.DB

// Function to open connection

func Connect() {
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successful connection")
	db = conection
}

// Function to close connection

func Close() {
	db.Close()
}

// Function to verify connection

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

type Object struct {
	Id int
	Image string
	Department string
	Location string
	Price int
	Area int
	Creation string
}

func Query() []Object {
	rows, err := db.Query("SELECT * FROM properties")
	if err != nil {
		panic(err.Error())
	}

	resp := Object{}
	arrayObject := []Object{}

	for rows.Next() {
		var id, price, area int
		var image, department, location, creation string

		err = rows.Scan(&id,&image,&department,&location,&price,&area,&creation)

		if err!= nil {
			panic(err.Error())
		}
		resp.Id = id
		resp.Image = image
		resp.Department = department
		resp.Location = location
		resp.Price = price
		resp.Area = area
		resp.Creation = creation

		arrayObject = append(arrayObject, resp)
	}

	return arrayObject
}