package summary

import (
	"fmt"
	"net/url"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/kkdai/youtube/v2"
)

type summaryService struct {
	summaryRepository ISummaryRepository
	youtubeClient     *youtube.Client
}

// CreateFromYoutubeVideo implements ISummaryService.
func (s *summaryService) CreateFromYoutubeVideo(userId string, videoUrl string) (*Summary, error) {
	// Download youtube video
	// Store it in Supabase storage

	videoId, err := getYoutubeVideoId(videoUrl)
	if err != nil {
		return nil, err
	}

	video, err := s.youtubeClient.GetVideo(videoId)
	if err != nil {
		return nil, err
	}
	formats := video.Formats.WithAudioChannels()
	fmt.Println(formats)

	return nil, nil
}

func NewSummaryService(summaryRepository ISummaryRepository, youtubeClient *youtube.Client) ISummaryService {
	return &summaryService{
		summaryRepository: summaryRepository,
	}
}

func getYoutubeVideoId(urlStr string) (string, error) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	query := parsedUrl.Query()
	videoId := query.Get("v")
	if videoId == "" {
		return "", custom_error.NewBadRequestError(fmt.Sprintf("no video ID found in URL: %s", urlStr))
	}
	return videoId, nil
}
