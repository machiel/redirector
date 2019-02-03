package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	listen := os.Getenv("LISTEN")
	redirectTo := os.Getenv("REDIRECT_TO")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectTo, http.StatusFound)
	})
	fmt.Printf("Ready to redirect to %s on %s\n", redirectTo, listen)
	http.ListenAndServe(listen, nil)
}
