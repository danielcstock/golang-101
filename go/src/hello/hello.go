package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			printLogs()
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
	fmt.Println("Monitorando...")

	sites := getSitesNamesFromFile()

	for _, site := range sites {
		time.Sleep(delay)
		siteTesting(site)
	}
}

func getSitesNamesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo.", err)
		return nil
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		sites = append(sites, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func siteTesting(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao monitorar o site.", err)
	} else if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso.")
	} else {
		fmt.Println("Site", site, "está com problemas. Status code", response.Status)
	}
	createLog(site, response.StatusCode == 200)

}

func createLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo.", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Status:" + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	fmt.Println("Exibindo logs...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo.", err)
	}

	fmt.Println(string(file))
}
