package middleware

import "net/http"

func DefaultPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello this is https"))
}
