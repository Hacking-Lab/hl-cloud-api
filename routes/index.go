package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Hacking-Lab/hl-cloud-api/core"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	output, _ := json.Marshal(map[string]interface{}{
		"message": "Welcome to the HL Cloud API!",
		"version": "v" + core.Version,
		"endpoints": map[string]string{
			"info":            "/info",
			"license_checker": "/license/check",
			"pricing":         "/pricing",
		},
	})
	w.Write(output)
}
