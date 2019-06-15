package healthcheck

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthCheck - basic model for a healthcheck
type HealthCheck struct {
	Msg     string `json:"msg"`
	Version string `json:"version"`
}

// Get - Simple Get of application healthcheck
func Get(w http.ResponseWriter, req *http.Request) {
	healthCheck := HealthCheck{"OK", "local"}
	data, err := json.Marshal(healthCheck)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
