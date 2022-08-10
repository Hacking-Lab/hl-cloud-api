package main

import (
	"fmt"
	"net/http"
)

func helloWorld(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "This should work!")
}

func main() {
	http.HandleFunc("/test", helloWorld)
	http.ListenAndServe(":1337", nil)
}
