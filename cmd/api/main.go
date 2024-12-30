package main

import (
	"github.com/gin-gonic/gin"
	"go_graphql_gin/internal/adapters/graphql"
	"go_graphql_gin/internal/adapters/notification"
	"go_graphql_gin/internal/adapters/repository/mongodb"
	"go_graphql_gin/internal/config"
	"go_graphql_gin/internal/core/services"
	"go_graphql_gin/internal/middleware"
	PkgMongodb "go_graphql_gin/pkg/mongodb"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB client
	mongoClient, err := PkgMongodb.NewClient(cfg.MongoDB)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	userRepo := mongodb.NewUserRepository(mongoClient)
	profileRepo := mongodb.NewProfileRepository(mongoClient)

	// Initialize notification service
	notificationSvc := notification.NewHTTPNotificationService(config.NotificationService.URL)

	// Initialize business services
	userService := services.NewUserService(userRepo, profileRepo, notificationSvc)

	// Initialize GraphQL handler
	graphqlHandler := graph.NewHandler(userService)

	// Initialize and start server
	router := gin.Default()

	// Add Auth0 middleware
	router.Use(middleware.AuthMiddleware(cfg.Auth0.Domain, cfg.Auth0.Audience))

	// Setup routes
	router.POST("/graphql", graphqlHandler.GraphQLHandler())
	router.GET("/playground", graphqlHandler.PlaygroundHandler())

	router.Run(":8080")
}
