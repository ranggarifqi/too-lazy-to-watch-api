package v1_summary

type CreateFromYoutubeDTO struct {
	Url string `json:"url" validate:"required"`
}
