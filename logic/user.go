package logic

import (
	"bblog/dao/mysql"
	"bblog/models"

	"bblog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// exist -> uid -> user insert into mysql
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	userId := snowflake.GenID()

	user := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}

	return mysql.InsertUser(user)
}
