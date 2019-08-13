package main

import (
  "<%= baseName %>/controllers"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
)

func main() {
  router := httprouter.New()
  router.GET("/", controllers.Hello)
  log.Fatal(http.ListenAndServe(":9001", router))
}
