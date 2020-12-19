package http

import (
	"github.com/gin-gonic/gin"
	"github.com/protonhq/proton/delivery/http/graphql"
	"github.com/protonhq/proton/registry"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// NewRouter - Create Service Routes
func NewRouter(ctn *registry.Container) *gin.Engine {
	router := gin.New()
	log := logrus.New()

	router.Use(ginlogrus.Logger(log))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/graphql", graphql.Handler(ctn))
	router.GET("/graphql", graphql.Handler(ctn))
	return router
}
