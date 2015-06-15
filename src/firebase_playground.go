package firebase_playground

import (
	"net/http"
)

func init() {
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// some code
}
