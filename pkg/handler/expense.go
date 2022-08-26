package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/g14com0/go-app/pkg/data"
	"github.com/gorilla/mux"
)

type Expense struct {
	l *log.Logger
}

func NewExpense(l *log.Logger) *Expense {
	return &Expense{l}
}

func (e *Expense) GetExpenses(rw http.ResponseWriter, r *http.Request) {
	e.l.Println("getexpenses")

	le := data.GetExpense()

	err := le.ToJSON(rw)

	if err != nil {
		http.Error(rw, "error", http.StatusInternalServerError)
	}
}

func (e *Expense) AddExpense(rw http.ResponseWriter, r *http.Request) {
	e.l.Println("addExpense")

	expense := r.Context().Value(KeyExpense{}).(*data.Expense)
	data.AddExpense(expense)
}

func (e *Expense) UpdateExpense(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "error", http.StatusBadRequest)
		return
	}

	e.l.Println("updateExpense")
	expense := r.Context().Value(KeyExpense{}).(*data.Expense)

	err = data.UpdateExpense(id, expense)
	if err == data.ErrExpenseNotFound {
		http.Error(rw, "error", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "error", http.StatusInternalServerError)
		return
	}
}

type KeyExpense struct{}

func (e *Expense) MiddlewareExpenseValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		expense := &data.Expense{}

		err := expense.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}

		err = expense.Validate()

		if err != nil {
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyExpense{}, expense)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
