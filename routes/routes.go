package routes

import (
	productData "my-project-be/features/product/data"
	productHandler "my-project-be/features/product/handler"
	productServices "my-project-be/features/product/services"

	userData "my-project-be/features/user/data"
	userHandler "my-project-be/features/user/handler"
	userServices "my-project-be/features/user/services"
	"my-project-be/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(c *echo.Echo, db *gorm.DB) {
	userData := userData.NewModel(db)
	userService := userServices.NewService(userData)
	userHandler := userHandler.NewUserHandler(userService)

	productData := productData.ProductModel(db)
	productService := productServices.ProductService(productData)
	productHandler := productHandler.ProductHandler(productService)

	c.POST("/register", userHandler.Register)
	c.POST("/login", userHandler.Login)
	c.GET("/keeplogin", userHandler.KeepLogin,middlewares.JWTMiddleware()) 
	c.PATCH("/update", userHandler.Update,middlewares.JWTMiddleware())

	c.POST("/product", productHandler.CreateProduct,middlewares.JWTMiddleware())
	c.GET("/search", productHandler.GetAllProduct)
	c.GET("/product/:productID", productHandler.GetProductById)
}