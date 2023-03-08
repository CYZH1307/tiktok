package controller

import (
	"github.com/CYZH1307/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	daoUser, _ := service.ParseToken(token)

	if daoUser.Id == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videoSvc := service.VideoSvc{}

	err := videoSvc(c, daoUser.Id, title)

	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Upload error"})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  title + " 已发布",
	})
}

func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")

	id, _ := strconv.ParseInt(userId, 10, 64)
	videoSvc := service.VideoSvc 
	myUser, _ := service.ParseToken(token)
	daoVideos := videoSvc.GetVideoListByUser(id)

	
	var videos []Video 
	for _, daoVideo := range daoVideos {
		video := ConvertVideo(&daoVideo, myUser.Id)
		videos = append(videos, video)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
