package gateway

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/repository"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service/service_model"
	"net/http"
	"strconv"
)

var Validate = validator.New()

type accountHandler struct {
	accountService service.Account
}

type accountPayload struct {
	Owner    string `json:"owner" validate:"required,min=2,max=10"`
	Balance  int64  `json:"balance" validate:"required,min=10"`
	Currency string `json:"currency" validate:"required,min=3"`
}

func (a *accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		badRequestResponse(w, r, err)
		return
	}

	account, err := a.accountService.GetAccount(r.Context(), id)
	if err != nil {
		switch err {
		case repository.ErrNotFound:
			notFoundResponse(w, r, err)
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}
	if err := jsonResponse(w, http.StatusOK, account); err != nil {
		internalServerError(w, r, err)
	}
}

func (a *accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	p := service_model.Pagination{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	pq, err := p.Parse(r)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(pq); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	accounts, err := a.accountService.GetAccounts(r.Context(), pq)
	if err != nil {
		internalServerError(w, r, err)
		return
	}
	if err := jsonResponse(w, http.StatusOK, accounts); err != nil {
		internalServerError(w, r, err)
	}
}

func (a *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var payload accountPayload
	if err := readJSON(w, r, &payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	account := &service_model.Account{
		Owner:    payload.Owner,
		Balance:  payload.Balance,
		Currency: payload.Currency,
	}

	if err := a.accountService.CreateAccount(r.Context(), account); err != nil {
		internalServerError(w, r, err)
		return
	}

	if err := jsonResponse(w, http.StatusCreated, account); err != nil {
		internalServerError(w, r, err)
		return
	}

}

func (a *accountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var payload accountPayload

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		badRequestResponse(w, r, err)
		return
	}

	account, err := a.accountService.GetAccount(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			notFoundResponse(w, r, err)
			return
		default:
			internalServerError(w, r, err)
			return
		}
	}

	if err := readJSON(w, r, &payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		badRequestResponse(w, r, err)
		return
	}

	account.Owner = payload.Owner
	account.Balance = payload.Balance
	account.Currency = payload.Currency

	if err := a.accountService.UpdateAccount(r.Context(), account); err != nil {
		internalServerError(w, r, err)
		return
	}

	if err := jsonResponse(w, http.StatusOK, account); err != nil {
		internalServerError(w, r, err)
		return
	}

}

func (a *accountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	if err := a.accountService.DeleteAccount(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			notFoundResponse(w, r, err)
		default:
			internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewAccountHandler(accountService service.Account) *accountHandler {
	return &accountHandler{
		accountService: accountService,
	}
}
