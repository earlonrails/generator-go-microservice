package main

import (
  "<%= baseName %>/controllers"
  "github.com/labstack/echo"
  mw "github.com/labstack/echo/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(mw.Logger())
  e.Use(mw.Recover())

  e.Get("/", controllers.Hello)

  // Start server
  e.Run(":9001")
}
