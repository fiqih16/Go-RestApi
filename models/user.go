package models

type User struct {
	ID int64 `gorm:"primary_key" json:"id"`
	Username string `gorm:"varchar(25)" json:"username"`
	Email string `gorm:"varchar(50)" json:"email"`
	Password string `gorm:"varchar(255)" json:"password"`
}

func IsEmailExist(email string) bool {
	var user User
	DB.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}