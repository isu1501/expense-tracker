package services

import (
	"errors"
	"expense-tracker/src/model"
	"expense-tracker/src/repository"
	"fmt"
	"strconv"
	"time"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(description string, amount string) error {

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return errors.New("invalid amount")
	}
	if description == "" || amountInt <= 0 {
		return errors.New("invalid expensedata")
	}

	expenses, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	exp := model.Expense{
		Id:          repository.GetNextId(expenses),
		Description: description,
		Amount:      amountInt,
		Created_at:  time.Now(),
	}

	expenses = append(expenses, exp)

	err = s.repo.Add(expenses)
	if err != nil {
		return err
	} else {
		fmt.Printf("Expense added successfully (Id: %d)\n", exp.Id)
	}

	return nil
}

func (s *ExpenseService) ListExpenses() error {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}

	fmt.Println("# ID  Date       Description  Amount")
	fmt.Println("--------------------------------------")
	for _, expense := range expenses {
		fmt.Printf("# %-3d %-12s %-14s $%-10d\n", expense.Id, expense.Created_at.Format("2006-01-02"), expense.Description, expense.Amount)
	}
	return nil
}

func (s *ExpenseService) SummaryExpense() error {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}

	err = s.repo.GetSummary(expenses)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExpenseService) SummaryExpenseByMonth(month string) error {
	expenses, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return errors.New("invalid month")
	}
	if monthInt < 1 || monthInt > 12 {
		return errors.New("invalid month")
	}

	err = s.repo.GetSummaryByMonth(monthInt, expenses)
	if err != nil {
		return err
	}
	return nil
}

func (s *ExpenseService) DeleteExpense(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid id")
	}

	expenses, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	err = s.repo.Delete(idInt, expenses)
	if err != nil {
		return err
	}
	fmt.Printf("Expense deleted successfully (Id: %d)\n", idInt)
	return nil
}
