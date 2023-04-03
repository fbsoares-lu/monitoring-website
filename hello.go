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

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readFile()

	// ForEach
	for _, site := range sites {
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("- Website: " + site + " - is working")
			registerLog(site, true)
		} else {
			fmt.Println("- Website: " + site + " - is not working")
			registerLog(site, false)
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("")
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showMenu() {
	fmt.Println("---------------------")
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Close the program")
	fmt.Println("---------------------")
}

func readCommand() int {
	var command int
	fmt.Scanf("%d", &command)

	return command
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Showing the logs...")
	fmt.Println(string(file))
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
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

func main() {
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Closing the program...")
			os.Exit(0)
		default:
			fmt.Println("Command does not exist")
			os.Exit(-1)
		}
	}

}
