package main

import (
	"my-project-be/config"
	"my-project-be/lib/redis"
	"my-project-be/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)
	rdb := redis.RedisClient(&cfg)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, db, rdb)
	e.Logger.Fatal(e.Start(":1300"))
}