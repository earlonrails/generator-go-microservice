package main

import (
  "<%= baseName %>/controllers"
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/", controllers.Hello)
  err := http.ListenAndServe(":9001", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
