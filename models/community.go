package models

import "time"

type Community struct {
	ID   int64  `gorm:"community_id"`
	Name string `gorm:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" gorm:"community_id"`
	Name         string    `json:"name" gorm:"community_name"`
	Introduction string    `json:"introduction,omitempty" gorm:"introduction"`
	CreateTime   time.Time `json:"create_time" gorm:"create_time"`
}
