package models

import (
	"database/sql"
	"time"
)

type Expense struct {
	ID          int
	Title       string
	Amount      float32
	Created     time.Time
	Category_id int
}

type ExpenseModel struct {
	DB *sql.DB
}
