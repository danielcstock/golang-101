package main

import "fmt"

func main() {
	showIntro()

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibier logs")

	var comando int
	fmt.Scan(&comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Comando desconhecido")
	}

	fmt.Println("Finalizando execução...")
}

func showIntro() {
	name := "Stock"
	version := 1.1
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}
