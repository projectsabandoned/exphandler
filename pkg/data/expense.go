package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Expense struct {
	ID          int     `json:"id"`
	Portfolio   string  `json:"portfolio"`
	Category    string  `json:"category"`
	SubCategory string  `json:"subcategory"`
	Import      float32 `json:"import"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (e *Expense) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(e)
}

type Expenses []*Expense

func (e *Expenses) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

func GetExpense() Expenses {
	return expenseList
}

func AddExpense(e *Expense) {
	e.ID = getNextID()
	expenseList = append(expenseList, e)
}

func UpdateExpense(id int, e *Expense) error {
	_, position, err := findExpense(id)
	if err != nil {
		return err
	}
	e.ID = id
	expenseList[position] = e

	return nil
}

var ErrExpenseNotFound = fmt.Errorf("expense not found")

func findExpense(id int) (*Expense, int, error) {
	for i, e := range expenseList {
		if e.ID == id {
			return e, i, nil
		}
	}
	return nil, -1, ErrExpenseNotFound
}

func getNextID() int {
	le := expenseList[len(expenseList)-1]
	return le.ID + 1
}

var expenseList = []*Expense{
	&Expense{
		ID:          1,
		Portfolio:   "Cash",
		Category:    "Food",
		SubCategory: "Bar",
		Import:      3.00,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Expense{
		ID:          2,
		Portfolio:   "Bank",
		Category:    "Shopping",
		SubCategory: "Pets",
		Import:      45.30,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
