package router

import (
	"expense-tracker/src/constants"
	"expense-tracker/src/repository"
	"expense-tracker/src/services"
)

func SetupRoutes(args []string) {

	repo := repository.NewExpenseRepository(constants.ExpensesFilePath)
	service := services.NewExpenseService(*repo)

	switch args[1] {
	case constants.AddCommand:
		service.AddExpense(args[3], args[5])
	case constants.ListCommand:
		service.ListExpenses()
	case constants.SummaryCommand:
		if len(args) < 3 {
			service.SummaryExpense()
		} else {
			service.SummaryExpenseByMonth(args[3])
		}
	case constants.DeleteCommand:
		service.DeleteExpense(args[3])
	}
}
