package logic

import (
	"bblog/dao/mysql"
	"bblog/models"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	return
}
