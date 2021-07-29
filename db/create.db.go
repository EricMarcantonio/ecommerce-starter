package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/products"
)

func CreateProduct(aProduct products.Product) (products.Product, error) {
	var err error
	var sqlQuery string
	var results *sql.Row
	var returnedProduct products.Product
	sqlQuery = fmt.Sprintf("insert into products(picName, description, price, takenBy) values ('%s', '%s', %f, '%s')", aProduct.Name, aProduct.Desc, aProduct.Price, aProduct.TakenBy)
	results = DB.QueryRow(sqlQuery)
	returnedProduct, err = ExtractProductFromRow(results)
	if err != nil {
		return products.Product{}, nil
	}
	return returnedProduct, nil
}
