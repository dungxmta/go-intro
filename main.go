package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/duzgxbe/go-intro/helper"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/demo", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8090"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func demo(c echo.Context) error {
	s := helper.Demo()
	return c.String(http.StatusOK, fmt.Sprintf("%v", s))
}
