package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xpanvictor/g3_template.git/internal/graph"
)

func SetupRouter(resolver *graph.Resolver) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(GinContextToContextMiddleware())

	// GraphQL endpoints
	r.POST("/graphql", graphqlHandler(resolver))
	if gin.Mode() != gin.ReleaseMode {
		r.GET("/playground", playgroundHandler())
	}

	return r
}
