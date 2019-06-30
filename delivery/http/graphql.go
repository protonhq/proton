package http

import (
	"github.com/protonhq/proton/delivery/http/schema"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// GraphqlHandler - GraphqlHandler handle graphql requests
func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
