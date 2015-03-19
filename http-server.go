package main

import("fmt"
"net/http")

// CREATE A HTTP SERVER

// http.ResponseWriter assembles the servers response and writes to
// the client
// http.Request is the clients request
func main() {
  
  // Calls for function handlers output to match the directory /
  http.HandleFunc("/", handler)
  http.HandleFunc("/zim", handlerZim)
  http.HandleFunc("/gir", handlerGir)

  // Listen to port 8080 and handle requests
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  // Writes to the client
  fmt.Fprintf(w, "Hello Earthlings!!!\n")
}

func handlerZim(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "DOOOOOOM!!!\n")
}

func handlerGir(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "doomdy doomdy doomdy doom ....\n")
}
