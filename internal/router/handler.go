package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xpanvictor/g3_template.git/internal/graph"
)

func SetupRouter(resolver *graph.Resolver) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(GinContextToContextMiddleware())

	// health check
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "alive"})
	})

	// GraphQL endpoints
	r.POST("/graphql", graphqlHandler(resolver))
	if gin.Mode() != gin.ReleaseMode {
		r.GET("/playground", playgroundHandler())
	}

	return r
}
