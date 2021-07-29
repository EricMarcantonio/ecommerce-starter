package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"graphql-go-pic-it/db"
	"graphql-go-pic-it/types"
	"log"
	"net/http"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    types.QueryType,
		Mutation: types.MutationType,
	},
)

func main() {
	var err = db.CreatePoolConn()
	db.Seed()

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/", h)
	log.Println("Listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("error listening on port " + err.Error())
	}
}