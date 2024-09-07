package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithOutProg := os.Args[1:]

	if len(argsWithOutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithOutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	maxConInput, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Errorf("Error - strconv.Atoi: %v", err)
	}
	maxPagesInput, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Errorf("Error - strconv.Atoi: %v", err)
	}
	fmt.Println("starting crawl")
	fmt.Println(rawBaseURL)

	maxConcurrency := maxConInput
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPagesInput)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL, maxPagesInput)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)

}
