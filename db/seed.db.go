package db

import (
	"context"
	"github.com/jackc/pgconn"
	"graphql-go-pic-it/products"
)

func CreateTableProducts() (pgconn.CommandTag, error) {
	const query = `
	create table products
	(
		id serial not null,
		"picName" varchar(128) not null,
		description varchar(512) not null,
		price float not null,
		"takenBy" varchar(64) not null,
		primary key (id)
	);`
	res, err := GlobalPoolConnection.Exec(context.Background(), query)
	if err != nil {
		return nil, err
	} else {
		return res, err
	}
}

func Seed()  {

	_, err := CreateTableProducts()
	if err != nil {
		return
	}

	err = CreateProduct(products.Product{
		Name:    "Eric",
		Desc:    "Woop",
		Price:   4,
		TakenBy: "MM",
	})
	if err != nil {
		return
	}


}

