package main

import (
	"fmt"
	"github.com/zerbitx/subby/api/v1alpha1"
)

func main() {
	p := &v1alpha1.Person{Name: "ryan"}

	fmt.Println(p.Name)
}
