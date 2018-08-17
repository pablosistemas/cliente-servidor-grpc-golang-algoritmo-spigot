package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func RecuperaProgresso(rw http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(rw, "200, %q", html.EscapeString(r.URL.Path))
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/status", RecuperaProgresso).Methods("POST")

  log.Fatal(http.ListenAndServe(":7778", nil))
}