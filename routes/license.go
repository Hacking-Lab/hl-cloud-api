package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Hacking-Lab/hl-cloud-api/checkers"
)

func HandleLicense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		output, _ := json.Marshal(map[string]interface{}{
			"error": map[string]string{
				"code":    "400",
				"message": "URL parameter 'key' is missing!",
			},
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	keyProvided := keys[0]
	isValid := checkers.Licence(keyProvided)

	output, _ := json.Marshal(map[string]bool{
		"valid": isValid,
	})
	w.Write(output)
}
