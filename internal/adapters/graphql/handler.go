package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go_graphql_gin/internal/adapters/graphql/generated"
	GraphResolver "go_graphql_gin/internal/adapters/graphql/resolvers"
	"go_graphql_gin/internal/core/ports"
)

type Handler struct {
	graphqlHandler    *handler.Server
	playgroundHandler gin.HandlerFunc
}

type Config struct {
	UserService ports.UserService
	// Add other services as needed
}

func NewHandler(cfg Config) *Handler {
	// Create a new resolver with dependencies
	resolverCfg := GraphResolver.Config{
		UserService: cfg.UserService,
	}
	r := GraphResolver.NewResolver(resolverCfg)

	// Create GraphQL server config
	c := graph.Config{
		Resolvers: r,
	}

	// Create the GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	// Create the playground
	playground := playground.Handler("GraphQL", "/query")

	return &Handler{
		graphqlHandler:    srv,
		playgroundHandler: gin.WrapH(playground),
	}
}

func (h *Handler) GraphQLHandler() gin.HandlerFunc {
	return gin.WrapH(h.graphqlHandler)
}

func (h *Handler) PlaygroundHandler() gin.HandlerFunc {
	return h.playgroundHandler
}
