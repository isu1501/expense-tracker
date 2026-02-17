package repository

import (
	"encoding/json"
	"expense-tracker/src/model"
	"fmt"
	"os"
	"time"
)

type ExpenseRepository struct {
	filePath string
}

func NewExpenseRepository(filePath string) *ExpenseRepository {
	return &ExpenseRepository{filePath: filePath}
}

func (r *ExpenseRepository) GetAll() ([]model.Expense, error) {
	var expenses []model.Expense

	file, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Expense{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &expenses)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetNextId(exp []model.Expense) int {

	maxId := 0
	if len(exp) == 0 {
		return maxId + 1
	}

	for _, expense := range exp {
		if expense.Id > maxId {
			maxId = expense.Id
		}
	}
	return maxId + 1
}

func (r *ExpenseRepository) Add(expenses []model.Expense) error {

	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(expenses)
}

func (r *ExpenseRepository) Delete(id int, expenses []model.Expense) error {

	found := false

	for i := range expenses {
		if expenses[i].Id == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Expense not found")
		return nil
	}

	err := r.Add(expenses)
	if err != nil {
		return err
	}
	return nil

}

func (r *ExpenseRepository) GetSummary(expenses []model.Expense) error {
	total := 0
	for _, expense := range expenses {
		total += expense.Amount
	}
	fmt.Printf("Total expenses: $%d\n", total)
	return nil
}

func (r *ExpenseRepository) GetSummaryByMonth(month int, expenses []model.Expense) error {
	total := 0
	currentYear := time.Now().Year()
	for _, expense := range expenses {
		if expense.Created_at.Year() != currentYear {
			continue
		}
		if expense.Created_at.Month() == time.Month(month) {
			total += expense.Amount
		}
	}
	fmt.Printf("Total expenses for %s: $%d\n", time.Month(month), total)
	return nil
}
