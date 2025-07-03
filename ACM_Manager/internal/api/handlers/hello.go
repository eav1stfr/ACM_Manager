package handlers

import (
	"acmmanager/utils"
	"encoding/json"
	"net/http"
)

func PingTheServerToAvoidColdSleep(w http.ResponseWriter, r *http.Request) {
	message := "Server was pinged!"

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(message)

	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
		return
	}
}
