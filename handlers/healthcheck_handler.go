package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthcheckStatus struct {
	Status string `json:"status"`
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	var healthcheck = HealthcheckStatus{Status: "WORKING"}
	json.NewEncoder(w).Encode(healthcheck)
}
