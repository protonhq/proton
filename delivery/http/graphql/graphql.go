package graphql

import (
	"github.com/protonhq/proton/delivery/http/graphql/schema"
	"github.com/protonhq/proton/registry"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// Handler - GraphqlHandler handle graphql requests
func Handler(ctn *registry.Container) gin.HandlerFunc {
	schema, err := schema.Schema(ctn)
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
