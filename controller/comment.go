package controller

import (
	"fmt"
	"github.com/CYZH1307/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"` // json?,omitempty?
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	daoUser, _ := service.ParseToken(token)

	if daoUser.Id == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User does't exist"})
		return
	}

	actionType := c.Query("action_type")
	commentSvc := service.CommentSvc{}

	if actionType == "1" {
		videoId, _  := strconv.PraseInt(c.Query("video_id"), 10, 64)
		text := c.Query("comment_text")
		data := commentSvc.CommentNew(daoUser.Id, videoId, text)

		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0}, 
			Comment: Comment {
				Id: 			data.Id
				Content:		text
				User:			ConvertUser(&daoUser, daoUser.Id)
				CreateDate:		fmt.Sprintf("%d-%d", time.Now().Month(), time.Now().Day())
			}, 
		})

		return 
	} else {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

		svc.CommentDelete(commentId)
	}
}

func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	commentSvc := service.CommentSvc{}
	data := commentSvc.CommentList(videoId)

	var results []CommentList(videoId)
	for _, v := range data {
		results = append(results, ConvertComment(&v))
	}

	c.JSON(http.StatusOK, CommentListResponse {
		Response: 		Response{StatusCode: 0},
		CommentList: 	reuslts
	})
}




