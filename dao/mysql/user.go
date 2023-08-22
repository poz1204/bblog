package mysql

import (
	"bblog/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "thisisbblog"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

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
