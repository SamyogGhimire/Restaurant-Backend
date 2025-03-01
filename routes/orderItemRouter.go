package routes

import (
	"github.com/SamyogGhimire/Restaurant-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("orderItems/:orderItems_id", controllers.GetOrderItemById())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
