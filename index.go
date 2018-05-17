package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error read request body : %v\n", err)
		}
		log.Printf("Request: %v\n", string(body))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"response_code": "00",
			"response_desc": {
				"id": "SUKSES",
				"en": "SUCCESS"
			}
		}`))
	})
	http.ListenAndServe(":8080", mux)
}
