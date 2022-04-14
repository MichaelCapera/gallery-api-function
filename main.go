package main

import (
	"encoding/json"
	"fmt"
	"gallery/db"
	"net/http"
)

func main() {

	// DB Connection

	db.Connect()
	properties := db.Query()
	fmt.Println("Server on port 8080")
	db.Close()

	// Get Gallery

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Error try with GET Method")
			return

		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(properties)
	})

	server := http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
