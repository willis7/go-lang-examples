package main

import("fmt"
"net/http")

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/zim", handlerZim)
  http.HandleFunc("/gir", handlerGir)

  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Earthlings!!!\n")
}

func handlerZim(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "DOOOOOOM!!!\n")
}

func handlerGir(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "doomdy doomdy doomdy doom ....\n")
}
