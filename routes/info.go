package routes

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/Hacking-Lab/hl-cloud-api/core"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	output, err := json.MarshalIndent(map[string]interface{}{
		"author": core.Author,
		"name":   core.Name,
		"runtime": map[string]string{
			"architecture":     runtime.GOARCH,
			"version":          runtime.Version(),
			"operating_system": runtime.GOOS,
		},
		"version": core.Version,
	}, "", "\t")
	if err == nil {
		w.Write(output)
	}
}
