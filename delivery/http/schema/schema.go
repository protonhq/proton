package schema

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"account": &queryAccount,
	},
})

// Schema - graphql schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil,
})