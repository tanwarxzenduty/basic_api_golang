package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetProducts() []Product {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testing_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM product")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []Product{}
	for results.Next() {
		var prod Product
		// for each row, scan into the Product struct
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		products = append(products, prod)
	}

	return products

}

//note that in (code string) *Product, the code is the JSON one in Product struct
func GetProduct(code string) *Product {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testing_db")
	prod := &Product{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product where code=?", code)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return prod
}

func AddProduct(product Product) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testing_db")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO product (code,name,qty,last_updated) VALUES (?,?,?, now())",
		product.Code, product.Name, product.Qty)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
