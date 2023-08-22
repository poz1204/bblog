package controller

import (
	"bblog/dao/mysql"
	"bblog/logic"
	"bblog/models"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//获取参数和校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("sign up with error parama", zap.Error(err))

		// todo validator
	}
	//业务
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回
	ResponseSuccess(c, nil)
}
