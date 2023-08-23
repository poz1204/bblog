package mysql

import (
	"bblog/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	result := db.Table("community").Find(&communityList)

	if result.RowsAffected == 0 {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}
