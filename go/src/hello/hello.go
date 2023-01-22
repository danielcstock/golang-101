package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Stock"
	version := 1.1
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("O tipo da variável 'name' é:", reflect.TypeOf(name))
	fmt.Println("O tipo da variável 'version' é:", reflect.TypeOf(version))
}
