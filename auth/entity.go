package auth

import "time"

type Customers struct {
	ID 		 	string
	Username 	string	 
	Email	 	string	 
	Password 	string	
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}