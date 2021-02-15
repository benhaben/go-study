package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"my name"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")
		fmt.Println("info=", info)
		fmt.Println("doc=", doc)
	}
}
func main() {
	var r resume
	findTag(&r)
}
