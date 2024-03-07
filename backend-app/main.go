package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "Hello from Backend!"}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)

	// Ajoutez les options du middleware CORS ici
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	http.Handle("/", corsHandler(r))

	fmt.Println("Server is listening on port 80")
	http.ListenAndServe(":80", nil)
}
