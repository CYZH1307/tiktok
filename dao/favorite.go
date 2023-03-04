package dao

type Favorite struct {
	Id				int64
	UserId			int64
	VideoId			int64
}

func InsertFavorite(data *Favorite) error {
	err := DB.Create(data).Error
	Handle(err)
	return err
}

// 未处理err
func DeleteFavorite(data *Favorite) {
	DB.Where("user_id = ? AND video_id = ?", data.UserId, data.VideoId).Delete(data)
}

func GetFavoriteByUser(userId int64) ([]Favorite, error) {
	var resluts []Favorite
	err := DB.Where("user_id = ?", userId).Find(&resluts).Error
	Handle(err)
	return resluts, err
}

func GetFavoriteByVideo(videoId int64) ([]Favorite, error) {
	var resluts []Favorite
	err := DB.error("video_id = ?", videoId).Find(&resluts).Error

	Handle(err)
	return resluts, err
}