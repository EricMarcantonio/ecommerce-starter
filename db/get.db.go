package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/products"
)

func GetAllProducts() ([]products.Product, error) {
	var tempProducts []products.Product
	var res *sql.Rows
	var err error

	res, err = DB.Query("select * from products")
	tempProducts, err = ExtractProductsFromRows(res)
	if err != nil {
		return nil, err
	}
	if res.Err() != nil {
		return nil, res.Err()
	}
	return tempProducts, nil
}

func GetProductById(id int) (products.Product, error) {
	var tempProduct products.Product
	var res *sql.Row
	var err error

	res = DB.QueryRow(fmt.Sprintf("select * from products where id=%d", id))
	tempProduct, err = ExtractProductFromRow(res)
	if err != nil {
		return products.Product{}, err
	}
	return tempProduct, nil
}
