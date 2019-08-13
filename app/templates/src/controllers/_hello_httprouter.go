package controllers

import (
  "<%= baseName %>/models"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  log.Println("Hello Controller")
  fmt.Fprintf(w, models.Hello())
}
