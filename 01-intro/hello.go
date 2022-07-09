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

const monitoring = 3
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()

		input := readInput()

		// sempre tem que colocar uma condição que retorne apenas true ou false (== ou != por exemplo), não aceita if comando {}, apenas se "comando" fosse boolean.
		/* 	if comando == 1 {
			   fmt.Println("Monitorando...")
		   } else if comando == 2 {
			   fmt.Println("Exibindo logs...")
		   } else if comando == 0 {
			   fmt.Println("Saindo do programa")
		   } else {
			   fmt.Println("Não conheço este comando")
		   } */

		switch input {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Displaying logs...")
			printLogs()
		case 0:
			fmt.Println("Exiting application...")
			os.Exit(0)
		default:
			fmt.Println("Unknown input")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Homero" // posso declarar com := sem precisar usar var e o tipo
	version := 1.1
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This application is running version:", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit Application")
}

func readInput() int {
	var selectedInput int // posso declarar usando var e o tipo
	// fmt.Scanf("%d", &comando) // %d representa um inteiro - & é o endereço de memoria da variavel que quero salvar (ponteiro)
	fmt.Scan(&selectedInput) // Scan não preciso passar o tipo de dado, a variavel ja inferiu
	fmt.Println("Selected input", selectedInput)
	fmt.Println("")

	return selectedInput
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br/", "https://www.caelum.com.br/"}
	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "is having problems. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Println(string(file))
}
