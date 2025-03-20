package routes

import (
	orderData "my-project-be/features/order/data"
	orderHandler "my-project-be/features/order/handler"
	orderServices "my-project-be/features/order/services"

	cartData "my-project-be/features/cart/data"
	cartHandler "my-project-be/features/cart/handler"
	cartServices "my-project-be/features/cart/services"

	productData "my-project-be/features/product/data"
	productHandler "my-project-be/features/product/handler"
	productServices "my-project-be/features/product/services"

	userData "my-project-be/features/user/data"
	userHandler "my-project-be/features/user/handler"
	userServices "my-project-be/features/user/services"
	"my-project-be/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
)

func InitRoute(c *echo.Echo, db *gorm.DB, rdb *redis.Client, mc midtrans.Client) {

	orderData := orderData.OrderModel(db)
	orderService := orderServices.OrderService(orderData, mc)
	orderHandler := orderHandler.OrderHandler(orderService)


	productData := productData.ProductModel(db)
	productService := productServices.ProductService(productData)
	productHandler := productHandler.ProductHandler(productService)

	cartData := cartData.CartModel(rdb)
	cartService := cartServices.CartService(cartData)
	cartHandler := cartHandler.CartHandler(cartService)

	userData := userData.NewModel(db)
	userService := userServices.NewService(userData, cartData)
	userHandler := userHandler.NewUserHandler(userService)

	c.POST("/register", userHandler.Register)
	c.POST("/login", userHandler.Login)
	c.GET("/keeplogin", userHandler.KeepLogin,middlewares.JWTMiddleware()) 
	c.PATCH("/update", userHandler.Update,middlewares.JWTMiddleware())

	c.POST("/product", productHandler.CreateProduct,middlewares.JWTMiddleware())
	c.GET("/search", productHandler.GetAllProduct)
	c.GET("/product/:productID", productHandler.GetProductById)

	c.POST("/cart", cartHandler.AddCart,middlewares.JWTMiddleware())
	c.GET("/cart", cartHandler.GetCart,middlewares.JWTMiddleware())
	c.DELETE("/cart/:productID", cartHandler.DeleteCartByID,middlewares.JWTMiddleware())
	c.DELETE("/cart", cartHandler.DeleteCart,middlewares.JWTMiddleware())

	c.POST("/order", orderHandler.CreateOrder,middlewares.JWTMiddleware())
}