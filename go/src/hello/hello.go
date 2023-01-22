package main

import "fmt"

func main() {
	name := "Stock"
	version := 1.1
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibier logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)

	switch comando {
	case 1:
		fmt.Println("Monitoramento")
		break
	case 2:
		fmt.Println("Logs")
		break
	case 0:
		fmt.Println("Sair")
		break
	}
}
