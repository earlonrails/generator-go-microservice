package controllers

import (
  "<%= baseName %>/models"
  "net/http"
  "log"
  "io"
)

func Hello(w http.ResponseWriter, req *http.Request) {
  log.Println("Hello Controller")
  io.WriteString(w, models.Hello())
}
