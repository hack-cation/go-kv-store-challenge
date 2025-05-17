package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{key}", getKey)
	mux.HandleFunc("POST /{key}", updateKey)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

var store = make(map[string]string)

type Response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func getKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	val := store[key]

	resp := Response{
		Key:   key,
		Value: val,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
}

type Request struct {
	Value string `json:"value"`
}

func updateKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		return
	}

	store[key] = req.Value

	w.WriteHeader(http.StatusCreated)
}
