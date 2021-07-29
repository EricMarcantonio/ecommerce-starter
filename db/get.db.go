package db

import (
	"context"
	"fmt"
	"graphql-go-pic-it/products"
)

func GetAllProducts() ([]products.Product, error) {
	var tempProducts []products.Product
	var res, err = GlobalPoolConnection.Query(context.Background(), "select * from products")
	if err != nil {
		return nil, err
	}
	if res.Err() != nil {
		return nil, res.Err()
	}
	for res.Next() {
		var id int
		var picName string
		var desc string
		var price float64
		var takenBy string
		err = res.Scan(&id, &picName, &desc, &price, &takenBy)
		if err != nil {
			return nil, err
		}
		tempProducts = append(tempProducts, products.Product{
			ID:      id,
			Name:    picName,
			Desc:    desc,
			Price:   price,
			TakenBy: takenBy,
		})
	}
	return tempProducts, nil
}

func GetProductById(id int) ([]products.Product,error) {
	var tempProducts []products.Product
	var res, err = GlobalPoolConnection.Query(context.Background(), fmt.Sprintf("select * from products where id=%d", id))
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var id2 int
		var picName string
		var desc string
		var price float64
		var takenBy string
		err = res.Scan(&id2, &picName, &desc, &price, &takenBy)
		if err != nil {
			return nil, err
		}
		tempProducts = append(tempProducts, products.Product{
			ID:      id2,
			Name:    picName,
			Desc:    desc,
			Price:   price,
			TakenBy: takenBy,
		})
	}
	return tempProducts, nil
}
