package mysql

import "errors"

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
