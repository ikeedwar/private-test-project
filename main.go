package main

import (
	"github.com/ikeedwar/private-test-project/rw"
	"github.com/ikeedwar/private-test-project/rwtest"
	"time"
)

func init() {
	println("main init")
}

func main() {

	var rwr rw.Reader = &rwtest.RwStruct{S: "test2", I: 9}

	var b string = "test"

	var c *string = &b

	var i, e = rwr.RwRead(c)

	println(i, " and ", e, " and ", *c)

	var t1 = rwtest.TestStruct{H: "h"}

	rwtest.PointTest("g", &t1)

	println(t1.H)

	for true {
		println("test")
		time.Sleep(time.Duration(2) * time.Second)
	}
}
