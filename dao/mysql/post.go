package mysql

import (
	"bblog/models"
)

func CreatePost(p *models.Post) (err error) {
	result := db.Table("post").Create(&p)

	return result.Error
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	result := db.Table("post").Where("post_id = ?", pid).First(&post)
	err = result.Error

	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	var offset = int((page - 1) * size)
	var limit = int(size)
	result := db.Table("post").Select("post_id, title, content, author_id, community_id, create_time").
		Order("create_time desc").
		Limit(limit).
		Offset(offset).
		Find(&posts)

	return posts, result.Error
}
