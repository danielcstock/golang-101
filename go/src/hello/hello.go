package main

import (
	"fmt"
	"os"
)

func main() {
	showIntro()

	showMenu()

	command := scanCommand()
	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Comando desconhecido")
		os.Exit(-1)
	}

	fmt.Println("Finalizando execução...")
	os.Exit(0)
}

func showIntro() {
	name := "Stock"
	version := 1.1
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}

func scanCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibier logs")
}
