package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/src/products"
)

func CreateProduct(aProduct products.Product) (products.Product, error) {
	var err error
	var sqlQuery string
	var results *sql.Rows
	var returnedProduct products.Product
	sqlQuery = fmt.Sprintf("insert into products(picName, description, price, takenBy) values ('%s', '%s', %f, '%s')", aProduct.Name, aProduct.Desc, aProduct.Price, aProduct.TakenBy)
	results, _ = Query(sqlQuery)
	_, err = ExtractProductsFromRows(results)
	if err != nil {
		return products.Product{}, nil
	}
	return returnedProduct, nil
}
