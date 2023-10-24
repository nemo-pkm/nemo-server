package router

import (
	"github.com/gin-gonic/gin"
	"nemo/controllers"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"pong": "hello",
		})
	})
	user := r.Group("/user")
	{
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"userAdd": "add",
			})
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"userDelete": "hel",
			})
		})
		user.POST("/update", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"userUpdate": "uio",
			})
		})
		user.POST("/list", controllers.GetUserList)
		user.GET("/info/:id/:name", controllers.GetUserInfo)
		user.POST("/signup", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"userSingUp": "use",
			})
		})
	}
	return r
}
