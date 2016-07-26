package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("o", "page.html", "Filename to use for saving response.")

	flag.Parse()

	url := os.Args[len(os.Args)-1]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("%v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	ioutil.WriteFile(*fileName, b, 0644)

}
