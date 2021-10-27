package main

import (
	"os"

	halloweenproduct "github.com/RN0311/demo-project/domain/halloweenProduct"
	halloweenproduct_handlers "github.com/RN0311/demo-project/handlers/halloweenProduct_handlers"
	"github.com/RN0311/demo-project/mongo"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := os.Getenv("MONGO_ADDRESS")

	mongoClient, err := mongo.NewClient(conn)
	if err != nil {
		panic("Error while connecting")
	}
	repo := halloweenproduct.NewRepository(mongoClient)

	engine := gin.New()
	engine.Use(gin.Recovery())

	halloweenproduct_handlers.InitializeRoutes(engine, repo)

	engine.Run(":8090")
}
