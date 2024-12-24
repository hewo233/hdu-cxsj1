package module

type User struct {
	Uid      int    `json:"uid" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique;size:100;not null"`
	Email    string `json:"email" gorm:"unique;size:100;not null"`
	Password string `json:"password" gorm:"size:255;not null" `
	Gender   string `json:"gender" gorm:"size:10"`
}

func NewUser() *User {
	return &User{}
}
