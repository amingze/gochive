package main

import (
	"fmt"
	"net/http"
)

const servePort = 6000

func main() {
	fmt.Println("service run in:", servePort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", servePort), nil)
	if err != nil {
		fmt.Println(err)
	}
}
