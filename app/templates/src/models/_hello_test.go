package models

import (
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
  Convey("Given hello model", t, func() {
    Convey("Hello should return 'Hello World!\n'", func() {
      So(Hello(), ShouldEqual, "Hello, World!\n")
    })
  })
}
