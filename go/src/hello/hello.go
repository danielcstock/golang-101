package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 5 * time.Second

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
	fmt.Println("2 - Exibir logs")
}
func startMonitoring() {
	sites := getSitesNames()

	fmt.Println("Monitorando...")
	for _, site := range sites {
		time.Sleep(delay)
		siteTesting(site)
	}
}

func getSitesNames() []string {
	sites := []string{"https://random-status-code.herokuapp.com/", "https://wwww.alura.com.br"}

	return sites
}

func siteTesting(site string) {
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site", site, "está com problemas. Status code", response.Status)
	}
}
