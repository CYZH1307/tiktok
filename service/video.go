package service

type VideoSvc struct {

}

func (vs *VideoSvc) GetVideoById(id int64) dao.Video {
	video, err := dao.GetVideoById(id)
	Handle(err)
	return video
}

func (vs *VideoSvc) GetVideoListByUser(userId int64) []dao.Video {
	videos, err := dao.GetVideoListByUser(userId)
	Handle(err)
	return videos
}

func (vs *VideoSvc) GetVideoByTime(time time.Time) []dao.Video {
	videos, err := dao.GetvideoListByTime(time)
	Handle(err)
	return videos
}

func (vs *videoSvc) SaveVideo(c *gin.Context, userId int64, title string) error {
	videoName, videoType, err := vs.getVideoName(c, userId)
	coverName := vs.getCoverName(videoNamem videoType)
	video := dao.Video {
		UserId:			userId,
		Title:			title,
		VideoUrl:		videoName + videoType,
		CoverUrl: 		coverName,
		PublishTime:	time.Now()
	}
	err := dao.insertVideo(&video)
	Handle(err)
	return err
}

func (vs * videoSvc) getVideoName(c *gin.Context, userId int64) (string, string, error) {
	data, err := c.FormFlie("data")
	Handle(err)
	filename := filepath.Base(data.Filename)
	ext := filepath.Ext(filename)
	key := config.Secret
	name, err := Encrypt(strconv.FormatInt(userId, 10), key)
	Handle(err)
	path := fmt.Sprintf(config.Video["video_dir_fmt"], name, ext)
	err := dao.SaveVideo(c, data, path)
	Handle(err)
	return name, ext, err
}

func (vs *videoSvc) getCoverName(filename, filetype string) string {
	name := dao.SaveCover(filename, filetype)
	return name
}