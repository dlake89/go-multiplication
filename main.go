package main
import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", Greet)
    log.Println("Starting server :8080")
    s := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
    }
    if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal("Server startup failed")
    }
}
func Greet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}
