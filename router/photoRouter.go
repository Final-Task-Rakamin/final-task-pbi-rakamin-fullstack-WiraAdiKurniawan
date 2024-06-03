package router

import (
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/controllers"
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/middlewares"
	"github.com/gin-gonic/gin"
)

func PhotoRoutes(r *gin.Engine) {
	photoGroup := r.Group("/photo")
	{
		photoGroup.GET("/", middlewares.Auth, controllers.GetPhoto)
		photoGroup.POST("/", middlewares.Auth, controllers.AddPhoto)
		photoGroup.PUT("/:id", middlewares.Auth, controllers.UpdatePhoto)
		photoGroup.DELETE("/:id", middlewares.Auth, controllers.DeletePhoto)
	}
}
