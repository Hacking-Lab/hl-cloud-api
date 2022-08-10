package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Hacking-Lab/hl-cloud-api/core"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	output, err := json.MarshalIndent(map[string]interface{}{
		"message": "Welcome to the HL Cloud API!",
		"version": "v" + core.Version,
		"endpoints": map[string]string{
			"info":            "/info",
			"license_checker": "/license/check",
			"pricing":         "/pricing",
		},
	}, "", "\t")
	if err == nil {
		w.Write(output)
	}
}
