package types

import (
	"github.com/graphql-go/graphql"
	"graphql-go-pic-it/src/resolvers"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolvers.GetProductById,
			},
			"products": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve:     resolvers.ListAllProducts,
			},
		},
	})
