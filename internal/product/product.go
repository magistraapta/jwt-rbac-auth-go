package product

import (
	"rbac-go/internal/product/controllers"
	"rbac-go/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.Engine, db *gorm.DB) {
	productController := controllers.NewProductController(db)

	router.POST("/products", middleware.RoleMiddleware("admin"), productController.CreateProduct)
	router.GET("/products", productController.GetAllProducts)

}
