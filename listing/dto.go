package listing

import "time"

type RequestProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
}

type ResponseProduct struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Quantity    string    `json:"quantity"`
	CreatedAt   time.Time `json:"created"`
	UpdateAt    time.Time `json:"updated"`
}
