package entity

import "time"

type Transaction struct {
	Id        int
	Amount    int
	status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
