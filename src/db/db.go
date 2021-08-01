package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/src/products"
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

	log.Printf("Trying to connect to %s", connectionString)
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	for DB.Ping() != nil {
		log.Println(fmt.Sprintf("DB is not ready. Waiting 4 seconds up to a max of %d", timeout))
		timeout = timeout - 4
		if timeout < 0 {
			panic("Couldn't connect in time. Please increase timeout")
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

func ExtractProductsFromRows(rows *sql.Rows) ([]products.Product, error) {
	var tempProducts []products.Product
	var colNames []string
	var err error
	colNames, err = rows.Columns()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tempProduct products.Product

		cols := make([]interface{}, len(colNames))
		for i := 0; i < len(colNames); i++ {
			cols[i] = ProductCol(colNames[i], &tempProduct)
		}
		err := rows.Scan(cols...)
		if err != nil {
			return nil, err
		}
		tempProducts = append(tempProducts, tempProduct)
	}
	if rows.Err() != nil {
		panic(err)
	}
	return tempProducts, nil
}

func ProductCol(colname string, product *products.Product) interface{} {
	switch colname {
	case "id":
		return &product.ID
	case "picName":
		return &product.Name
	case "description":
		return &product.Desc
	case "price":
		return &product.Price
	case "takenBy":
		return &product.TakenBy
	default:
		panic("Not impletmented")
	}
}

func Query(sql string) (*sql.Rows, error) {
	log.Printf("Querying DB with the string: %s", sql)
	query, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	return query, nil

}

func QueryRow(sql string) *sql.Row {
	log.Printf("Querying DB with the string: %s", sql)
	return DB.QueryRow(sql)
}

func Exec(sql string) (sql.Result, error) {
	log.Printf("Exec on DB with the string: %s", sql)
	return DB.Exec(sql)

}
