package routes

import (
	"encoding/json"
	"net/http"
)

func GetPricing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	output, err := json.MarshalIndent(map[string]interface{}{
		"plans": map[string]string{
			"Coming": "soon",
		},
	}, "", "\t")
	if err == nil {
		w.Write(output)
	}
}
