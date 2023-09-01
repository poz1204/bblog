package mysql

import (
	"bblog/models"
	"crypto/md5"
	"encoding/hex"
)

const secret = "thisisbblog"

func CheckUserExist(username string) (err error) {
	var count int64
	result := db.Table("user").Where("username = ?", username).Select("COUNT(*)").Scan(&count)

	if result.Error != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	result := db.Table("user").Create(&user)
	return result.Error
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	thisPass := user.Password
	mysqlUser := models.User{}

	result := db.Table("user").First(&mysqlUser)
	if result.RowsAffected == 0 {
		return ErrorUserNotExist
	}
	password := mysqlUser.Password
	if password != thisPass {
		return ErrorInvalidPassword
	}

	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	result := db.Table("user").Where("user_id = ?", uid).First(&user)
	return user, result.Error
}
