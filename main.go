package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	httpPort = 6000
)

func init() {
	flag.IntVar(&httpPort, "port", 6000, "p")
}

func main() {
	fmt.Println("service run in:", httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
	if err != nil {
		fmt.Println(err)
	}
}
