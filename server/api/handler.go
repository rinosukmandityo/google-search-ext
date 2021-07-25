package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PostSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Only accept POST request")
		http.Error(w, "Only accept POST request", http.StatusBadRequest)
		return
	}

	data, err := getBodyData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("Search Result:")
	fmt.Println(string(data))
	w.Write(data)
}

func getBodyData(r *http.Request) (res []byte, err error) {
	decoder := json.NewDecoder(r.Body)
	respBody := make([]Search, 0)
	if err = decoder.Decode(&respBody); err != nil {
		log.Println("cannot decode response body", err.Error())
		return
	}

	res, err = json.MarshalIndent(respBody, "", "\t")
	if err != nil {
		log.Println("cannot Marshal response", err.Error())
		return
	}

	return
}
