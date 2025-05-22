package auth

import (
	"rbac-go/internal/auth/controllers"
	"rbac-go/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	router.POST("/login", authController.Login)

	group := router.Group("/admin")
	{
		group.POST("/register", authController.RegisterAdmin)
		group.GET("/dashboard", middleware.RoleMiddleware("admin"), func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "admin dashboard",
			})
		})
	}

	group = router.Group("/user")
	{
		group.POST("/register", authController.RegisterUser)
		group.GET("/dashboard", middleware.RoleMiddleware("user"), func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "user dashboard",
			})
		})
	}
}
