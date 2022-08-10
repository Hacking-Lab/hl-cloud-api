package main

import (
	"net/http"

	"github.com/Hacking-Lab/hl-cloud-api/routes"
)

func main() {
	http.HandleFunc("/info", routes.GetInfo)
	http.HandleFunc("/license/check", routes.CheckLicense)
	http.HandleFunc("/pricing", routes.GetPricing)
	http.ListenAndServe(":1337", nil)
}
