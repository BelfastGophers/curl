package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("o", "", "Filename to use for saving response.")
	httpMethod := flag.String("X", "GET", "HTTP method to use for the request")
	reqBody := flag.String("d", "", "Request body")

	flag.Parse()

	url := os.Args[len(os.Args)-1]

	var err error
	var resp *http.Response

	if *httpMethod == "GET" {
		resp, err = http.Get(url)
	} else {
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(*reqBody)))
		req.Header.Set("Content-type", "application/json")

		client := &http.Client{}

		resp, err = client.Do(req)
	}

	if err != nil {
		fmt.Println("%v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if *fileName == "" {
		fmt.Println(string(b))
	} else {
		ioutil.WriteFile(*fileName, b, 0644)
	}

}
