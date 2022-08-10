package main

import (
	"net/http"

	"github.com/Hacking-Lab/hl-cloud-api/routes"
)

func main() {
	http.HandleFunc("/info", routes.HandleInfo)
	http.HandleFunc("/license/check", routes.HandleLicense)
	http.HandleFunc("/pricing", routes.HandlePricing)
	http.ListenAndServe(":1337", nil)
}
