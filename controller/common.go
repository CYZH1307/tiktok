package controller

import (
	"fmt"
	"github.com/CYZH1307/tiktok/config"
	"github.com/CYZH1307/tiktok/dao"
	"github.com/CYZH1307/tiktok/service"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func ConvertVideo(video *dao.Video, myId int64) Video {
	var userSvc 		service.userSvc
	var commentSvc 		service.CommentSvc
	var favoriteSvc 	service.FavoriteSvc

	user := userSvc.GetUserById(video.UserId)
	videoPrefix := config.Video["video_prefix"]
	converPrefix := config.Video["conver_prefix"]

	return Video {
		Id: 			video.Id
		Author:			ConvertUser(&User, 0)
		Title: 			video.Title
		PlayUrl:       	videoPrefix + video.VideoUrl,
		CoverUrl:      	coverPrefix + video.CoverUrl,
		FavoriteCount: 	int64(len(favoriteSvc.FavoriteListByVideo(video.Id))),
		CommentCount: 	int64(len(commentSvc.CommentList(video.Id))),
		IsFavorite:    	favoriteSvc.IsFavorite(myId, video.Id),
	}
}

func ConvertUser(user *dao.User, myId int64) User {
	relationSvc := service.RelationSvc{
		MyId: myId,
	}
	return User{
		Id:            	user.Id,
		Name:          	user.Name,
		FollowCount:   	relationSvc.LenUserFocus(user.Id),
		FollowerCount: 	relationSvc.LenUserFans(user.Id),
		IsFollow:      	relationSvc.IsFollow(user.Id),
	}
}

func ConvertComment(comment *dao.Comment) Comment {
	var svc service.UserSvc
	user := svc.GetUserById(comment.UserId)
	return Comment{
		Id:         	comment.Id,
		User:       	ConvertUser(&user, 1),
		Content:    	comment.CommentText,
		CreateDate: 	fmt.Sprintf("%d-%d", comment.CreateDate.Month(), comment.CreateDate.Day()),
	}
}







