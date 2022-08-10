package main

import (
	"net/http"

	"github.com/Hacking-Lab/hl-cloud-api/routes"
)

func main() {
	http.HandleFunc("/info", routes.GetInfo)
	http.ListenAndServe(":1337", nil)
}
