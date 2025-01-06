package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xpanvictor/g3_template.git/bootstrap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Server init!")

	app := bootstrap.App()
	defer app.CleanUp()

	env := app.Env
	//db := app.Mongo.Database(env.DBName)

	// ctx management
	//timeout := time.Duration(env.ContextTimeout) * time.Second

	g := gin.Default()

	srv := &http.Server{
		Addr:    env.ServerAddress,
		Handler: g.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server listening e at : %s\n", err)
		}
	}()

	// run graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully shutting down server!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %s", err)
	}
	select {
	case <-ctx.Done():
		log.Print("Timeout of 5 seconds.")
	}
	log.Println("Server exiting.")

}
