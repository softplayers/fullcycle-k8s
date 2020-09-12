package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(s string) string {
	return "<b>" + s + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>%s</body></html>", greeting("Code.education Rocks!"))
}

func main() {
	fmt.Print("Listening on 8000...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
