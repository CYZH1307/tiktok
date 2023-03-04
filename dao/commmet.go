package dao

import "time"

type Comment struct {
	Id				int64
	UserId			int64
	VideoId			int64
	CommentText		string
	CreateDate		time.Time
}

func InsertComment(data *Comment) error {
	err := DB.Create(data).Error
	Handle(err)
	return err
}

// 未处理err
func DeleteComment(commentId int64) {
	data := Comment{Id, commentId};
	DB.Where("id = ?", commentId).Delete(&data)
}

func GetCommnet(videoId int64) ([]Comment, error) {
	var resluts []Comment
	err := DB.Where("video_id = ?", videoId).Find(&resluts).Error

	Handle(err)

	return resluts, err
}