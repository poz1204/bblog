package routes

import (
	"bblog/controller"
	"bblog/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")

	// 注册 登陆
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	{
		v1.GET("/community", controller.CommunityHandler)
		//v1.GET("/community/:id", controller.CommunityDetailHandler)

		//v1.POST("/post", controller.CreatePostHandler)
		//v1.GET("/post/:id", controller.GetPostDetailHandler)
		//v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		//v1.GET("/posts2", controller.GetPostListHandler2)

		// 投票
		//v1.POST("/vote", controller.PostVoteController)

	}

	return r
}
