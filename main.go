package main

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
 http.Redirect(w, r, "/", http.StatusFound)
}

