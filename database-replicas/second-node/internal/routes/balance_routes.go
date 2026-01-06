package routes

import (
	"second-node/internal/controller"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	balances := r.Group("/balances")
	{
		balances.POST("", controller.CreateBalance)
		balances.PUT("", controller.UpdateBalance)
		balances.GET("/:id", controller.GetBalance)
		balances.GET("", controller.GetAllBalances)
		balances.DELETE("/:id", controller.DeleteBalance)
	}
}
