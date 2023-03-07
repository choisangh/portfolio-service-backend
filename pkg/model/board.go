package model

type Board struct {
	ID          uint   `gorm:"primarykey;autoIncrement" json:"_id"`
	CompanyName string `json:"companyName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Location    string `json:"location"`
	Link        string `json:"link"`
	Description string `json:"description"`
}
