package listing

import "time"

type Products struct {
	Id          string
	Name        string
	Description string
	Quantity    int
	Price       float64
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
