package main

import (
	"fmt"
	"github.com/cts5ws/reprox"
	"net/http"
)

func routeRequest(res http.ResponseWriter, req *http.Request) {
	reprox.Route(res, req)
}

func main() {

	http.HandleFunc("/", routeRequest)

	fmt.Println("Starting proxy server on port 80")
	http.ListenAndServe(":80", nil)
}
