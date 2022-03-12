// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shodan/shodan"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("Usage: main <searchterm> <page> <page increment>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	var next string
	p_num, con_err := strconv.Atoi(os.Args[2])
	p_inc, p_err := strconv.Atoi(os.Args[3])
	if con_err != nil {
		log.Fatalln("Useage: Page must be a number")
	}
	if p_err != nil {
		log.Fatalln("Useages: Page Increment must be a number")
	}
	next = "y"
	for next == "y" {
		fmt.Printf(
			"Query Credits: %d\nScan Credits:  %d\n\n",
			info.QueryCredits,
			info.ScanCredits)

		hostSearch, err := s.HostSearch(os.Args[1], p_num)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("Host Data Dump\n")
		for _, host := range hostSearch.Matches {
			fmt.Println("==== start ", host.IPString, "====")
			h, _ := json.Marshal(host)
			fmt.Println(string(h))
			fmt.Println("==== end ", host.IPString, "====")

		}

		fmt.Printf("IP, Port\n")

		for _, host := range hostSearch.Matches {
			fmt.Printf("%s, %d\n", host.IPString, host.Port)
		}
		fmt.Println("Press y and Enter Key to get next page.")
		fmt.Scanln(&next)

		p_num += p_inc
	}

}
