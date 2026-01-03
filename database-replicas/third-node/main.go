package main

import (
	"os"
	"third-node/internal/database"
	"third-node/internal/initializer"
	"third-node/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initializer.InitializeEnvironment()
	database.ConnectDB()
	r := gin.Default()
	routes.GetRoutes(r)
	r.Run(":" + os.Getenv("REPLICA_PORT"))
}
