package main

import (
    "fmt"
    "net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "MarketAffiliateHelp Web — OK")
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "MarketAffiliateHelp — Starter Web Service Online")
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/health", health)

    fmt.Println("🚀 MAH Web Service running on :8080")
    http.ListenAndServe(":8080", nil)
}
