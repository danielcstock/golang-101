package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()

	for {
		showMenu()
		command := scanCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		default:
			fmt.Println("Finalizando execução...")
			os.Exit(0)
		}
	}

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
func startMonitoring() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site", site, "está com problemas. Status code", response.Status)
	}

}
