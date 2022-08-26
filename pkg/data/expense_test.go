package data

import "testing"

func TestCheckValidation(t *testing.T) {
	e := &Expense{
		Portfolio: "Cash",
		Category:  "bar",
	}
	err := e.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
