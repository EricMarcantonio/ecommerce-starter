package main

import (
	"fmt"
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
	var err = db.CreateConn()
	if err != nil {
		panic("Error creating connection to DB")
	}
	err = db.Seed()
	if err != nil {
		fmt.Println(err)
	}

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
