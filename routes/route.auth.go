package route

import "github.com/gin-gonic/gin"

func InitAuthRoutes(route *gin.Engine) {

	auth := route.Group("/api/v1/auth")

	auth.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(200, "hello")
	})
	auth.POST("/login")

}
