package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/protonhq/proton/domain"
)

var accountType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Account",
	Description: "Account Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"userName": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryAccount = graphql.Field{
	Name:        "QueryAccount",
	Description: "Query Account",
	Type:        graphql.NewList(accountType),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"userName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		name, _ := p.Args["userName"].(string)
		res := []domain.Account{
			{UserName: name},
		}
		return res, nil
	},
}
