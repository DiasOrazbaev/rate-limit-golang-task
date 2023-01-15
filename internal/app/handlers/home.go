package handlers

import "net/http"

func Home(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Home page"))
	if err != nil {
		return
	}
}
