package main

import (
	"fmt"
	"reflect"

	"github.com/kylelemons/godebug/pretty"
)

type foo struct {
	id   int
	name string
}

func main() {
	f1 := &foo{id: 1, name: "hoge"}
	f2 := &foo{id: 1, name: "hoge"}
	f3 := &foo{id: 2, name: "hoge"}
	if reflect.DeepEqual(f1, f2) {

	}
	fmt.Println("f1 and f2 %#v", reflect.DeepEqual(f1, f2))
	fmt.Println("f2 and f3 %#v", reflect.DeepEqual(f2, f3))
	fmt.Println("f1 and f3 %#v", reflect.DeepEqual(f1, f3))
	fmt.Println(pretty.Compare(f1, f3))
}
