package lookup

import (
	"encoding/json"
	"net/http"
)

// Handler exposes the lookup HTTP endpoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	msisdn := r.URL.Query().Get("msisdn")
	if msisdn == "" {
		http.Error(w, "missing msisdn parameter", http.StatusBadRequest)
		return
	}

	resp := LookupResponse{
		MSISDN:      msisdn,
		Country:     Country(msisdn),
		NumberType:  NumberType(msisdn),
		ValidLength: IsValidLength(msisdn),
		Operator:    Operator(msisdn),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
