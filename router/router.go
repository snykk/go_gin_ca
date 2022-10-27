package router

import (
	"github.com/gin-gonic/gin"

	"github.com/snykk/go_gin_ca/controllers/middleware"
	v1 "github.com/snykk/go_gin_ca/controllers/v1"
	"github.com/snykk/go_gin_ca/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin)
			auth.POST("/regis", v1.POSTRegister)
		}

		user := v1route.Group("/user")
		{
			user.GET("/:id", utils.AuthOnly, v1.GETUser)
			user.PUT("", utils.AuthOnly, v1.PUTUser)
		}

		todo := v1route.Group("/todo")
		{
			todo.GET("", utils.AuthOnly, v1.GETAllUserTodo)
			todo.GET("/:id", utils.AuthOnly, v1.GETTodoByID)
			todo.PUT("/:id", utils.AuthOnly, v1.PUTTodoUser)
			todo.POST("", utils.AuthOnly, v1.POSTTodo)
			todo.DELETE("/:id", utils.AuthOnly, v1.DELETETodo)
		}
	}
	return
}
