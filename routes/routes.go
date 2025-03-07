package routes

import (
	"my-project-be/features/user/data"
	"my-project-be/features/user/handler"
	"my-project-be/features/user/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(c *echo.Echo, db *gorm.DB) {
	userData := data.NewModel(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	c.POST("/register", userHandler.Register)
	c.POST("/login", userHandler.Login)
}