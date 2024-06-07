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

	videoPath, err := s.downloadYoutubeVideo(videoId, id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Video downloaded: %s\n", videoPath)

	// Upload it to Supabase storage
	uploadedUrl, err := s.summaryRepository.UploadVideo(videoPath, id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Uploaded %v\n", uploadedUrl)

	// Store in db
	summaryPayload := &CreateSummaryPayload{
		Id:              id,
		UserId:          userId,
		Status:          "PENDING",
		VideoUrl:        uploadedUrl,
		YoutubeVideoUrl: videoUrl,
	}
	summary, err := s.summaryRepository.Create(*summaryPayload)
	if err != nil {
		return nil, err
	}

	// Delete local tmp file
	if err = os.Remove(videoPath); err != nil {
		return nil, err
	}

	return summary, nil
}

const VIDEO_QUALITY string = "360"

func (s *summaryService) downloadYoutubeVideo(youtubeVideoId string, uniqueId string) (string, error) {
	video, err := s.youtubeClient.GetVideo(youtubeVideoId)
	if err != nil {
		return "", err
	}
	formats := video.Formats.WithAudioChannels().Quality(VIDEO_QUALITY)

	stream, _, err := s.youtubeClient.GetStream(video, &formats[0])
	if err != nil {
		return "", err
	}
	defer stream.Close()

	videoPath := fmt.Sprintf("./tmp/%v.mp4", uniqueId)

	file, err := os.Create(videoPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "", err
	}
	return videoPath, nil
}

func NewSummaryService(summaryRepository ISummaryRepository, youtubeClient youtube.Client) ISummaryService {
	return &summaryService{
		summaryRepository: summaryRepository,
		youtubeClient:     youtubeClient,
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
