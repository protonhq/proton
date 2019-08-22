package http

import (
	"github.com/protonhq/proton/delivery/http/schema"
	"github.com/protonhq/proton/registry"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// GraphqlHandler - GraphqlHandler handle graphql requests
func GraphqlHandler(ctn *registry.Container) gin.HandlerFunc {
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
