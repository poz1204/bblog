package mysql

import (
	"bblog/models"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	result := db.Table("community").Model(&models.Community{}).Select("community_id, community_name").Find(&communityList)
	if result.RowsAffected == 0 {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	result := db.Table("community").Where("community_id = ?", id).First(&community)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = ErrorInvalidID
		}
	}

	return community, err
}
