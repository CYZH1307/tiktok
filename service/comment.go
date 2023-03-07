package service

type CommentSvc struct {

}

func (cs *CommentSvc) CommentSvc(userId int64, videoId int64, commentText string) dao.Comment {
	commentTime := time.Now() // 能不能直接嵌入的下面？
	data := dao.Comment({UserId: userId, VideoId: videoId CommentText: commentText, CreateDate: commentTime})

	_ = dao.InsertComment(&data)
	return data
}

func (cs *CommentSvc) CommentDelete(commentId int64) {
	dao.DeleteComment(commentId)
}

func (cs *CommentSvc) CommentList(videoId int64) []dao.Comment {
	data, _ := dao.GetComment(videoId)

	return data
}