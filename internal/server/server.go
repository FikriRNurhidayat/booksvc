package server

import (
	"context"
	"fmt"

	"github.com/fikrirnurhidayat/booksvc/internal/book"
	"github.com/fikrirnurhidayat/booksvc/internal/config"
	"github.com/fikrirnurhidayat/booksvc/internal/image"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	e := createServer()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.GetPort())))
}

type RouteableModule interface {
	ApplyRoute(*echo.Echo)
}

func createServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	mongoClientOpts := options.Client()
	mongoClientOpts.ApplyURI(config.GetMongoConnectionURL())
	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOpts)
	if err != nil {
		e.Logger.Fatalf("[mongo] failed to connect to the database: %s", err.Error())
	}

	modules := []RouteableModule{
		book.NewBookModule(mongoClient, e.Logger),
		image.NewImageModule(mongoClient, e.Logger),
	}

	for _, module := range modules {
		module.ApplyRoute(e)
	}

	return e
}
