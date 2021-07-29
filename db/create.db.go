package db

import (
	"context"
	"fmt"
	"graphql-go-pic-it/products"
)

func CreateProduct(aProduct products.Product) error {
	var sqlQuery = fmt.Sprintf("insert into products (\"picName\", description, price, \"takenBy\") values ('%s', '%s', %f, '%s')", aProduct.Name, aProduct.Desc, aProduct.Price, aProduct.TakenBy)
	_, err := GlobalPoolConnection.Exec(context.Background(), sqlQuery)
	return err
}
