package server

import (
	"Echo-Simple-JWT/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// login route
	e.POST("/login", handler.Login())

	// restricted route
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/hello", handler.Restricted())
	e.Logger.Fatal(e.Start(":1323"))
}
