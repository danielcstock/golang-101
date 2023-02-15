package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rawLines := readFiles()

	writeFormattedFile(rawLines)
}

func readFiles() []string {
	var inputs []string
	file, err := os.Open("text_001")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo.", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, _ := reader.ReadString('\n')
		formattedLine := strings.TrimSpace(line)
		if formattedLine != "Game Log Output Begins Here:" && formattedLine != "Game Log Output Ends Here" {
			inputs = append(inputs, formattedLine)
		} else if formattedLine == "Game Log Output Ends Here" {
			break
		}
	}
	file.Close()

	return inputs
}

func writeFormattedFile(lines []string) {
	file, err := os.OpenFile("formattedFile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo.", err)
	}

	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	file.Close()
}
