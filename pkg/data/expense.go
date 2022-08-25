package data

import (
	"encoding/json"
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

type Expenses []*Expense

func (e *Expenses) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}

func GetExpense() Expenses {
	return expenseList
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
		ID:          1,
		Portfolio:   "Bank",
		Category:    "Shopping",
		SubCategory: "Pets",
		Import:      45.30,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
