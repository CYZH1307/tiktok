package dao

type Video struct {
	Id				int64
	UserId			int64
	Title			string
	VideoUrl		string
	CoverUrl		string
	PublishTime		time.Time
}

func GetVideoById(id int64) (Video, error) {
	var video Video
	err := DB.Where("id = ?", id).First(&video).Error
	Handle(err)
	return video, err
}

func GetVideoListByUser(userId int64) ([]Video, error) {
	var videos []Video
	err := DB.Where("user_id = ?", userId).First(&videos).Error
	Handle(err)

	return videos, err
}

func GetVideoListByTime(time, time.Time) ([]Video, error) {
	limit, _ := strconv.ParseInt(config.Video["limit"], 10, 64)
	videos := make([]Video, limit)
	err := DB.where("publish_time < ?", time).Order("publish_time desc").Limit(int(limit)).Find(&videos).Error
	Handle(err)

	return videos, err
}

func InsertVideo(video *Video) error {
	err := DB.Create(&video).Error
	Handle(err)
	return err
}

func SaveVideo(c *gin.Context, data *multipart.FileHeader, path string) error {
	err := c.SaveUploadedFile(data, path)
	Handle(err)
	return err
}

// SaveCover 保证视频方面
func SaveCover(filename string, filetype string) string {
	//格式化
	inputFile := fmt.Sprintf(config.Video["video_dir_fmt"], filename, filetype)
	outputDir := fmt.Sprintf(config.Video["cover_dir_fmt"], filename)
	// 
	cmd := exec.Command("./tools/ffmpeg", "-i", inputFile, "-vframes", "1", outputDir)
	err := cmd.Run() //阻塞Go程序，直到命令完成
	Handle(err)
	return filename + ".png"
}

