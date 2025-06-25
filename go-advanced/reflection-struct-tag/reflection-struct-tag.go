package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "ANkush", Age: 26}
	fmt.Println(p)

	t := reflect.TypeOf(p)
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag.Get("json"))
}
