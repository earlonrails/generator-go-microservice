package routes

import (
  "<%= baseName %>/models"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, models.Hello())
}
