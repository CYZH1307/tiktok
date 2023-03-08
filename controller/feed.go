package controller

import (
	"github.com/CYZH1307/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	token := c.Query("token")

	myUser, _ := service.ParseToken(token)
	var feedTime time.Time 
	if latestTime != "0" {
		latest, _ := strconv.ParseInt(latestTime, 10, 64)
		feedTime = time.Unix(latest, 0)
	} else {
		feedTime = time.Now()
	}

	videoSvc := service.VideoSvc{}

	daoVideos := videoSvc.GetVideoListByTime(time.Now())

	var videos []Video

	for _, daoVideo := range daoVideos {
		video := ConvertVideo(&daoVideo, myUser.Id)
		videos = append(videos, video)

	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  feedTime.Unix(),
	})
}









