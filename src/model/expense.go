package model

import "time"

type Expense struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Created_at  time.Time `json:"created_at"`
}
