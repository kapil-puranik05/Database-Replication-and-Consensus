package main

import (
	"os"
	"second-node/internal/database"
	"second-node/internal/initializer"
	"second-node/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initializer.InitializeEnvironment()
	database.ConnectDB()
	r := gin.Default()
	routes.GetRoutes(r)
	r.Run(":" + os.Getenv("REPLICA_PORT"))
}
