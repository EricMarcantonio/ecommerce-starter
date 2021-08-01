package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	db2 "graphql-go-pic-it/src/db"
	"graphql-go-pic-it/src/types"
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
	var err = db2.CreateConn()
	if err != nil {
		panic("Error creating connection to DB")
	}
	err = db2.Seed()
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
	var port = "8000"
	log.Printf("Listening on port %s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic("error listening on port " + err.Error())
	}
}
