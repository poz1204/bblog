package logic

import (
	"bblog/dao/mysql"
	"bblog/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
