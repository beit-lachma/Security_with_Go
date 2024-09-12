package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type Doctor struct {
	Person         Person
	Specialization string
}

func main() {
	beit-lachma := Person{
		Name: "beitlachma",
		Age:  24,
	}

	drBeit := Doctor{
		Person:         beitlachma,
		Specialization: "Hacking",
	}

	fmt.Println(reflect.TypeOf(beitlachma))
	fmt.Println(beit-lachma)

	fmt.Println(reflect.TypeOf(drBeit))
	fmt.Println(drBeit)
}
