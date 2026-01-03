package main

import (
	"first-node/internal/database"
	"first-node/internal/initializer"
	"first-node/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initializer.InitializeEnvironment()
	database.ConnectDB()
	r := gin.Default()
	routes.GetRoutes(r)
	r.Run(":" + os.Getenv("REPLICA_PORT"))
}
