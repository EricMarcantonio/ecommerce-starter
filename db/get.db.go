package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/products"
	"log"
	"strings"
)

func GetAllProducts(selectedFields []string) ([]products.Product, error) {
	var tempProducts []products.Product
	var res *sql.Rows
	var err error

	log.Println(fmt.Sprintf("select %s from products", BuildOptimizedSQL(selectedFields)))
	res, err = DB.Query(fmt.Sprintf("select %s from products", BuildOptimizedSQL(selectedFields)))

	tempProducts, err = ExtractProductsFromRows(res)
	if err != nil {
		return nil, err
	}
	if res.Err() != nil {
		return nil, res.Err()
	}
	return tempProducts, nil
}

func GetProductById(id int, requestedFields []string) (products.Product, error) {
	var tempProduct []products.Product
	var res *sql.Rows
	var err error

	res, err = DB.Query(fmt.Sprintf("select %s from products where id=%d", BuildOptimizedSQL(requestedFields), id))
	tempProduct, err = ExtractProductsFromRows(res)
	if err != nil {
		return products.Product{}, err
	}
	return tempProduct[0], nil
}

func BuildOptimizedSQL(requestedFields []string) string {
	var tempString strings.Builder
	var columnName []string
	for _, eachString := range requestedFields {
		if eachString == "id" {
			columnName = append(columnName, "id")
			continue
		}
		if eachString == "name" {
			columnName = append(columnName, "picName")
			continue
		}
		if eachString == "desc" {
			columnName = append(columnName, "description")
			continue
		}
		if eachString == "price" {
			columnName = append(columnName, "price")
			continue
		}
		if eachString == "takenBy" {
			columnName = append(columnName, "takenBy")
			continue
		}
	}
	if len(columnName) == 0 {
		panic("OOF")
	}
	if len(columnName) == 1 {
		tempString.WriteString(columnName[0])
	} else {
		for i, finalCol := range columnName {
			if i < len(columnName)-1 {
				tempString.WriteString(finalCol)
				tempString.WriteString(",")
			}
		}
		tempString.WriteString(columnName[len(columnName)-1])
	}
	return tempString.String()

}
