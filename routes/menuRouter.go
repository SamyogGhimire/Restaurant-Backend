package routes

import (
	"github.com/SamyogGhimire/Restaurant-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenuById())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
