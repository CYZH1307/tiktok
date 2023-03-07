package service

type FavoriteSvc struct {

}
// FavoriteAction 用户对视频执行点赞/取消点赞
func (fs *FavoriteSvc) FavoriteAction(userId, videoId int64, actionType int32) {
	dao := dao.Favorite{UserId: userId, VideoId: videoId}
	
	if actionType == 1 {
		_ = dao.InsertFavorite(&data)
	} else if actionType == 2 {
		dao.DeleteFavorite(&data)
	} else {
		log.Fatalln("action_type is wrong !")
	}

}
// FavoriteListByUser 通过用户得到关注的视频列表
func (fs *FavoriteSvc) FavoriteListByUser(userId int64) []dao.Video {
	favoriteVideos, _ := dao.GetFavoriteByUser(userId)

	Vsvc := VideoSvc{}
	var results []dao.Video

	for _, data := range favoriteVideos {
		result := Vsvc.GetVideoById(data.VideoId)
		results = append(results, result)
	}

	return results
}
// FavoriteListByVideo 通过视频得到关注的用户列表 
func (fs *FavoriteSvc) FavoriteListByVideo(videoId int64) []dao.user {
	favoriteUsers, _ := dao.GetFavoriteByVideo(videoId)

	Usvc := Usersvc{}
	var results []dao.User

	for _, data := range favoriteUsers {
		result := Usvc.GetUserById(data.UserId)
		results = append(results, result)
	}

	return results
}
// IsFavorite 判断用户是否点赞视频
func (fs *FavoriteSvc) IsFavorite(userId, videoId int64) bool {
	list, _ : = dao.GetFavoriteByUser(userId)
	result := false
	for _, data := range list {
		if data.VideoId == videoId {
			result = true
			break
		}

	}
	return result
}
