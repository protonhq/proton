package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/protonhq/proton/registry"
	"github.com/protonhq/proton/usecase"
)

// Schema - graphql schema
func Schema(ctn *registry.Container) (graphql.Schema, error) {
	// Queries
	accountUsecase := ctn.Resolve("account-usecase").(usecase.AccountUsecase)
	queryAcc := queryAccount(accountUsecase)

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Root Query",
		Fields: graphql.Fields{
			"account": &queryAcc,
		},
	})

	// Mutations
	createAcc := createAccount(accountUsecase)
	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootMutation",
		Description: "Root Mutation",
		Fields: graphql.Fields{
			"create_account": &createAcc,
		},
	})
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
