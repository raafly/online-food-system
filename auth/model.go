package auth

import "time"

type Customers struct {
	ID 		 	string
	Username 	string	 
	Email	 	string	 
	Password 	string	
	CreatedAt time.Time
	UpdatedAt time.Time
}