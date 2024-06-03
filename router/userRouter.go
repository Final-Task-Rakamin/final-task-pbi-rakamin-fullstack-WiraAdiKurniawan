package router

import (
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/controllers"
	"github.com/Final-Task-Rakamin/final-task-pbi-rakamin-fullstack-WiraAdiKurniawan/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.GET("/active", middlewares.Auth, controllers.GetUserLogin)
		userGroup.GET("/logout", controllers.Logout)

	}

	r.GET("/", controllers.PostsIndex)
}
