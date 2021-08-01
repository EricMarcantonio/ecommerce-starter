package db

import (
	"database/sql"
	"graphql-go-pic-it/src/products"
)

func CreateTableProducts() (sql.Result, error) {
	const query = `
	create table if not exists products 
(
	id int auto_increment,
	picName varchar(128) not null,
	description varchar(128) not null,
	price float not null,
	takenBy varchar(128) not null,
	constraint products_pk
		primary key (id)
);
`
	res, err := Exec("drop table if exists products;")
	res, err = Exec(query)
	if err != nil {
		return nil, err
	} else {
		return res, err
	}
}

func Seed() error {
	var err error
	_, err = CreateTableProducts()
	if err != nil {
		return err
	}
	_, err = CreateProduct(products.Product{
		Name:    "Eric Marcantonio",
		Desc:    "Picture of Eric Marcantonio",
		Price:   8.99,
		TakenBy: "MM",
	})
	_, err = CreateProduct(products.Product{
		Name:    "Michelle Mali",
		Desc:    "Picture of Michelle Mali",
		Price:   9.99,
		TakenBy: "Emarc",
	})
	return err
}
