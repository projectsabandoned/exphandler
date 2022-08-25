package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPost {
		e.addExpense(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "error", http.StatusBadRequest)
		}
		fmt.Println("got id: ", id)
		e.updateExpense(id, rw, r)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (e *Expense) getExpenses(rw http.ResponseWriter, r *http.Request) {
	e.l.Println("getexpenses")

	le := data.GetExpense()

	err := le.ToJSON(rw)

	if err != nil {
		http.Error(rw, "error", http.StatusInternalServerError)
	}
}

func (e *Expense) addExpense(rw http.ResponseWriter, r *http.Request) {
	e.l.Println("addExpense")

	expense := &data.Expense{}
	err := expense.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "error", http.StatusBadRequest)
	}
	data.AddExpense(expense)
}

func (e *Expense) updateExpense(id int, rw http.ResponseWriter, r *http.Request) {
	e.l.Println("updateExpense")

	expense := &data.Expense{}
	err := expense.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "error", http.StatusBadRequest)
	}
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
