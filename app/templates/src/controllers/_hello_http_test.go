package controllers

import (
  "net/http"
  "net/http/httptest"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
  Convey("Given hello controller", t, func() {
    Convey("Hello should return 'Hello World!\n'", func() {
      req, err := http.NewRequest("GET", "/", nil)
      w := httptest.NewRecorder()
      Hello(w, req)
      So(err, ShouldBeNil)
      So(w.Code, ShouldEqual, http.StatusOK)
      So(w.Body.String(), ShouldEqual, "Hello, World!\n")
    })
  })
}
