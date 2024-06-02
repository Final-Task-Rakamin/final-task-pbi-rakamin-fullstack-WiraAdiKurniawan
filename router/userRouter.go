package router

import (
	"github.com/adikrnwn171/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.GET("/active", controllers.GetUserLogin)
		userGroup.GET("/logout", controllers.Logout)

	}

	r.GET("/", controllers.PostsIndex)
}
