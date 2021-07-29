package resolvers

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"graphql-go-pic-it/db"
	"graphql-go-pic-it/products"
)

func GetProductById(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("use products for all products")
	}
	result, err := db.GetProductById(id, GetSelectedFields([]string{"product"}, p))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ListAllProducts(p graphql.ResolveParams) (interface{}, error) {
	var products2 []products.Product
	var err error
	products2, err = db.GetAllProducts(GetSelectedFields([]string{"products"}, p))
	return products2, err
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

func GetSelectedFields(selectionPath []string, resolveParams graphql.ResolveParams) []string {
	fields := resolveParams.Info.FieldASTs
	for _, propName := range selectionPath {
		found := false
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}
				found = true
				break
			}
		}
		if !found {
			return []string{}
		}
	}
	var collect []string
	for _, field := range fields {
		collect = append(collect, field.Name.Value)
	}
	return collect
}
