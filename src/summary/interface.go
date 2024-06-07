package summary

type CreateSummaryPayload struct {
	Id              string `json:"id"`
	UserId          string `json:"user_id"`
	Status          string `json:"status"`
	VideoUrl        string `json:"video_url"`
	YoutubeVideoUrl string `json:"youtube_video_url"`
}

const TABLE_NAME = "Summaries"

type ISummaryRepository interface {
	Create(payload CreateSummaryPayload) (*Summary, error)
	UploadVideo(videoPath string, uniqueId string) (string, error)
}

type ISummaryService interface {
	CreateFromYoutubeVideo(userId string, videoUrl string) (*Summary, error)
}
