package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID          uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Email       string `gorm:"uniqueIndex" json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Description string `gorm:"default:null" json:"description"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
