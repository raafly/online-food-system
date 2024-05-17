package helper

import "golang.org/x/crypto/bcrypt"

type Password struct{}

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if  err != nil {
		panic(err)
	}
	return string(hash)
}

func (p *Password) CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))	
}