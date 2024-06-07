package summary

import (
	"fmt"
	"io"
	"net/url"
	"os"
	custom_error "too-lazy-to-watch-api/src/error"

	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
)

type summaryService struct {
	summaryRepository ISummaryRepository
	youtubeClient     youtube.Client
}

// CreateFromYoutubeVideo implements ISummaryService.
func (s *summaryService) CreateFromYoutubeVideo(userId string, videoUrl string) (*Summary, error) {
	id := uuid.New().String()

	// Download youtube video
	videoId, err := getYoutubeVideoId(videoUrl)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Video ID: %s\n", videoId)

	video, err := s.youtubeClient.GetVideo(videoId)
	if err != nil {
		return nil, err
	}
	formats := video.Formats.WithAudioChannels().Quality("360")

	stream, _, err := s.youtubeClient.GetStream(video, &formats[0])
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	file, err := os.Create(fmt.Sprintf("./tmp/%v.mp4", id))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return nil, err
	}

	// Generate UUID
	// Store it in Supabase storage

	return nil, nil
}

func NewSummaryService(summaryRepository ISummaryRepository, youtubeClient youtube.Client) ISummaryService {
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
