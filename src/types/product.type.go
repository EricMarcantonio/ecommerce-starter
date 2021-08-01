package types

import "github.com/graphql-go/graphql"

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "ID of the image",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the image",
			},
			"desc": &graphql.Field{
				Type:        graphql.String,
				Description: "Description of the image",
			},
			"price": &graphql.Field{
				Type:        graphql.Float,
				Description: "Price of the Photo",
			},
			"takenBy": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the person who shot the picture",
			},
		},
	},
)
