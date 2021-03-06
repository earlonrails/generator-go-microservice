package controllers

import (
  "<%= baseName %>/models"
  "net/http"
  "github.com/labstack/echo"
  "log"
)

func Hello(c *echo.Context) error {
  log.Println("Hello Controller")
  return c.String(http.StatusOK, models.Hello())
}
