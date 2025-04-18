package routes

import (
	"github.com/SamyogGhimire/Restaurant-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFoods())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PUT("/foods/:food_id", controllers.UpdateFood())
}
