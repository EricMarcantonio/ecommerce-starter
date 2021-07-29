package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/products"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func CreateConn() error {
	var err error
	var host = os.Getenv("HOST")
	var port = os.Getenv("PORT")
	var db = os.Getenv("DB")
	var user = os.Getenv("USER")
	var pass = os.Getenv("PASS")
	var timeout, _ = strconv.Atoi(os.Getenv("TIMEOUT"))

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, db)

	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	for DB.Ping() != nil {
		log.Println(fmt.Sprintf("DB is not ready. Waiting 4 seconds up to a max of %d", timeout))
		timeout = timeout - 4
		if timeout < 0 {
			panic("Couldn't connect in time. Please increase the timeout in docker-compose")
		}
		time.Sleep(time.Second * 4)
		DB, err = sql.Open("mysql", connectionString)
		if err != nil {
			return err
		}
	}
	log.Println("Connection to DB made")

	return nil
}

func ExtractProductFromRow(row *sql.Row) (products.Product, error) {
	var tempProduct products.Product
	var err error
	err = row.Scan(&tempProduct.ID, &tempProduct.Name, &tempProduct.Desc, &tempProduct.Price, &tempProduct.TakenBy)
	if err != nil {
		return products.Product{}, err
	}
	return tempProduct, nil
}

func ExtractProductsFromRows(rows *sql.Rows) ([]products.Product, error) {
	var tempProducts []products.Product
	for rows.Next() {
		var tempProduct products.Product
		err := rows.Scan(&tempProduct.ID, &tempProduct.Name, &tempProduct.Desc, &tempProduct.Price, &tempProduct.TakenBy)
		if err != nil {
			return nil, err
		}
		tempProducts = append(tempProducts, tempProduct)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return tempProducts, nil
}
