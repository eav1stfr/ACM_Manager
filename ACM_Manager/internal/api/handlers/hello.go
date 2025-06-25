package handlers

import (
	"acmmanager/utils"
	"encoding/json"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello!"

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(message)

	if err != nil {
		http.Error(w, utils.EncodingResponseError.Error(), utils.EncodingResponseError.GetStatusCode())
		return
	}

	log.Println("Successfully responded to hello request")
}
