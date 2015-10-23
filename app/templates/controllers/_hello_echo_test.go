package controllers

import (
  "net/http"
  "net/http/httptest"
  "github.com/labstack/echo"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
  Convey("Given hello controller", t, func() {
    Convey("Hello should return 'Hello World!\n'", func() {
      e := echo.New()
      req, err := http.NewRequest(GET, "/", nil)
      w := httptest.NewRecorder()
      c := NewContext(req, NewResponse(rec), e)

      Hello(c)
      So(err, ShouldBeNil)
      So(w.Code, ShouldEqual, http.StatusOK)
      So(w.Body.String(), ShouldEqual, "Hello, World!\n")
    })
  })
}
