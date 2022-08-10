package routes

import (
	"encoding/json"
	"net/http"
)

func CheckLicense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	output, err := json.MarshalIndent(map[string]interface{}{
		"Hello": "world!",
	}, "", "\t")
	if err == nil {
		w.Write(output)
	}
}
