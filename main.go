package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)
	mux.HandleFunc("/multiply", Multiply)
	log.Println("Starting server :8080")
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var inputs Input
	err = json.Unmarshal([]byte(body), &inputs)

	inputs.Answer = inputs.A * inputs.B
	log.Println(inputs)

	response, err := json.Marshal(inputs)
	w.Write(response)
}

type Input struct {
	A      float32 `json:"a"`
	B      float32 `json:"b"`
	Answer float32 `json:"answer"`
}
