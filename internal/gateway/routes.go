package gateway

import "net/http"

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /accounts", nil)
	mux.HandleFunc("GET /accounts/{id}", nil)
	mux.HandleFunc("POST /accounts", nil)
	mux.HandleFunc("PUT /accounts/{id}", nil)
	mux.HandleFunc("DELETE /accounts/{id}", nil)

	return mux
}
