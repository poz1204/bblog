package logic

import (
	"bblog/dao/mysql"
	"bblog/models"
	"bblog/pkg/jwt"
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
func Login(p *models.ParamSignUp) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
