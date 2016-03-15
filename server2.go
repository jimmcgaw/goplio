package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
)

var mu sync.Mutex
var count int

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/count", counter)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(rw http.ResponseWriter, req *http.Request) {
  mu.Lock()
  fmt.Fprintf(rw, "Count %d\n", count)
  mu.Unlock()
}
