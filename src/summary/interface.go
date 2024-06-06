package summary

type CreateSummaryPayload struct {
	UserId   string `json:"user_id"`
	Status   string `json:"status"`
	VideoUrl string `json:"video_url"`
}

const TABLE_NAME = "Summaries"

type ISummaryRepository interface {
	Create(payload CreateSummaryPayload) (*Summary, error)
}

type ISummaryService interface {
	CreateFromYoutubeVideo(userId string, videoUrl string) (*Summary, error)
}
