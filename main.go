package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 || os.Args[1:][0] == "" {
		panic("No arguments specified")
	}

	uri := os.Args[1:][0]
	url, uriValidationError := url.ParseRequestURI(uri)

	if uriValidationError != nil {
		panic("Invalid url")
	}

	resp, err := http.Get(url.String())

	if err != nil {
		panic(fmt.Sprintf("Couldn't get results for : %s", url.String()))
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
