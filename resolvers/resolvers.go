package resolvers

import (
	"errors"
	"github.com/graphql-go/graphql"
	"graphql-go-pic-it/db"
	"graphql-go-pic-it/products"
)

func GetPictureById(p graphql.ResolveParams) (interface{}, error) {

	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("use list for all products")
	}
	result, err := db.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ListAllPictures(p graphql.ResolveParams) (interface{}, error) {
	var products []products.Product
	var err error
	products, err = db.GetAllProducts()
	return products, err
}

func CreateProduct(p graphql.ResolveParams) (interface{}, error) {
	product := products.Product{
		Name:    p.Args["name"].(string),
		Desc:    p.Args["desc"].(string),
		Price:   p.Args["price"].(float64),
		TakenBy: p.Args["takenBy"].(string),
	}
	product, err := db.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(p graphql.ResolveParams) (interface{}, error) {
	id, idOk := p.Args["id"].(int)
	name, nameOk := p.Args["name"].(string)
	desc, descOk := p.Args["info"].(string)
	price, priceOk := p.Args["price"].(float64)
	takenBy, takenByOk := p.Args["takenBy"].(string)

	if !idOk {
		return nil, errors.New("no id passed to UpdateProduct")
	}

	tempProduct := products.Product{ID: id}
	var res products.Product
	if nameOk {
		tempProduct.Name = name
	}
	if descOk {
		tempProduct.Desc = desc
	}
	if priceOk {
		tempProduct.Price = price
	}
	if takenByOk {
		tempProduct.TakenBy = takenBy
	}

	res, err := db.UpdateProduct(tempProduct)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteProduct(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	_ = db.DeleteProduct(id)
	return nil, nil
}
