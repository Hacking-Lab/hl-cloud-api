package routes

import (
	"encoding/json"
	"net/http"
)

func HandlePricing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	output, _ := json.Marshal(map[string]interface{}{
		"plans": map[string]interface{}{
			"free": map[string]string{
				"montly": "0 CHF",
				"yearly": "0 CHF",
			},
			"basic": map[string]string{
				"montly": "9.99 CHF",
				"yearly": "99.99 CHF",
			},
			"pro": map[string]string{
				"montly": "19.99 CHF",
				"yearly": "199.99 CHF",
			},
			"enterprise": map[string]string{
				"montly": "29.99 CHF",
				"yearly": "299.99 CHF",
			},
		},
	})
	w.Write(output)
}
