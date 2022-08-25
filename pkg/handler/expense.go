package handler

import (
	"log"
	"net/http"

	"github.com/g14com0/go-app/pkg/data"
)

type Expense struct {
	l *log.Logger
}

func NewExpense(l *log.Logger) *Expense {
	return &Expense{l}
}

func (e *Expense) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		e.getExpenses(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (e *Expense) getExpenses(rw http.ResponseWriter, r *http.Request) {
	le := data.GetExpense()
	//d, err := json.Marshal(le)
	err := le.ToJSON(rw)

	if err != nil {
		http.Error(rw, "error", http.StatusInternalServerError)
	}
}
