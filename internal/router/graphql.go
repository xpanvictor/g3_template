package router

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/xpanvictor/g3_template.git/internal/graph"
)

// Defining middleware for GraphQL
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GraphQL handler
func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers:  resolver,
		Directives: graph.DirectiveRoot{},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
