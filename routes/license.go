package routes

import (
	"encoding/json"
	"net/http"
)

func CheckLicense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		output, _ := json.MarshalIndent(map[string]interface{}{
			"error": map[string]string{
				"code":    "400",
				"message": "URL parameter 'key' is missing!",
			},
		}, "", "\t")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	key := keys[0]
	isValid := true
	if key != "5P3C1AL-4DM1N15TR4T0R-K3Y" {
		isValid = false
	}
	output, err := json.MarshalIndent(map[string]bool{
		"valid": isValid,
	}, "", "\t")
	if err == nil {
		w.Write(output)
	}
}
