package models

type User struct {
	UserId   int `gorm:"primaryKey;autoIncrement"`
	Mobile   string
	Name     string
	Password string
}

func CheckAuth(mobile string, password string) bool {
	var user User

	db.Select("user_id").Where(User{Mobile: mobile, Password: password}).First(&user)
	if user.UserId > 0 {
		return true
	}
	return false
}

func GetUser(mobile string, password string) User {
	user := User{}

	db.Where(User{Mobile: mobile, Password: password}).First(&user)
	return user
}
