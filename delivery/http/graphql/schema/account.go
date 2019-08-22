package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/protonhq/proton/domain"
	"github.com/protonhq/proton/usecase"
)

var accountType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Account",
	Description: "Account Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func queryAccount(uc usecase.AccountUsecase) graphql.Field {
	return graphql.Field{
		Name:        "QueryAccount",
		Description: "Query Account",
		Type:        graphql.NewList(accountType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			name, _ := p.Args["email"].(string)
			res := []domain.Account{
				{UserName: name},
			}
			return res, nil
		},
	}
}

func createAccount(uc usecase.AccountUsecase) graphql.Field {
	return graphql.Field{
		Type:        accountType,
		Description: "Create new account",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			email, _ := p.Args["email"].(string)
			password, _ := p.Args["password"].(string)

			acc, err := uc.RegisterUser(email, password)

			if err != nil {
				return nil, err
			}

			return acc, nil
		},
	}
}
