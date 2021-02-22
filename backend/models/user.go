package models

import "github.com/jinzhu/gorm"

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

// 手机号、密码查询用户
func GetUser(mobile string, password string) User {
	user := User{}

	db.Where(User{Mobile: mobile, Password: password}).First(&user)
	return user
}

// 判断手机号是否存在
func IsMobileExist(mobile string) (bool, error) {
	var user User
	err := db.Select("user_id").Where(User{Mobile: mobile}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.UserId > 0 {
		return true, nil
	}
	return false, nil
}

// 新增用户
func AddUser(mobile string, password string, name string) error {
	user := User{
		Mobile:   mobile,
		Name:     name,
		Password: password,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
