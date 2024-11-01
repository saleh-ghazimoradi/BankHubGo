package gateway

import "net/http"

type Handlers struct {
	GetAccount    http.HandlerFunc
	GetAccounts   http.HandlerFunc
	CreateAccount http.HandlerFunc
	UpdateAccount http.HandlerFunc
	DeleteAccount http.HandlerFunc
}

func Routes(handler Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /accounts", handler.GetAccounts)
	mux.HandleFunc("GET /accounts/{id}", handler.GetAccount)
	mux.HandleFunc("POST /accounts", handler.CreateAccount)
	mux.HandleFunc("PUT /accounts/{id}", handler.UpdateAccount)
	mux.HandleFunc("DELETE /accounts/{id}", handler.DeleteAccount)

	return mux
}
