package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestGreet tests the Greet function to ensure it returns the correct greeting.
func TestGreet(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Greet)

    handler.ServeHTTP(rr, req)

    expected := "Hello World!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

// TestMultiply tests the Multiply function to ensure it correctly multiplies two numbers and returns the result.
func TestMultiply(t *testing.T) {
    input := &Input{A: 2, B: 3}
    inputJSON, _ := json.Marshal(input)

    req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(inputJSON))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Multiply)

    handler.ServeHTTP(rr, req)

    var output Input
    _ = json.Unmarshal(rr.Body.Bytes(), &output)

    expected := float32(6)
    if output.Answer != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", output.Answer, expected)
    }
}