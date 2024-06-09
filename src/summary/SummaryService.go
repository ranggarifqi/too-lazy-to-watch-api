package summary

import (
	"fmt"
	"io"
	"net/url"
	"os"
	custom_error "too-lazy-to-watch-api/src/error"
	"too-lazy-to-watch-api/src/storage"
	"too-lazy-to-watch-api/src/taskPublisher"

	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
)

type summaryService struct {
	summaryRepository       ISummaryRepository
	youtubeClient           youtube.Client
	taskPublisherRepository taskPublisher.ITaskPublisherRepository
	storageRepository       storage.IStorageRepository
}

const TASK_CHANNEL = "summarization"
const BUCKET_NAME = "video"

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
	videoFile, err := os.Open(videoPath)
	if err != nil {
		return nil, err
	}
	defer videoFile.Close()

	cloudRelativePath := fmt.Sprintf("%s.mp4", id)
	uploadedUrl, err := s.storageRepository.Upload(BUCKET_NAME, cloudRelativePath, videoFile, storage.FileOptions{
		ContentType: "video/mp4",
	})
	if err != nil {
		deleteLocalTmpVideo(videoPath)
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
		deleteLocalTmpVideo(videoPath)
		s.storageRepository.DeleteFile(BUCKET_NAME, cloudRelativePath)
		return nil, err
	}

	if err = s.taskPublisherRepository.Publish(TASK_CHANNEL, taskPublisher.PublishPayload{
		ContentType: "text/plain",
		Body:        []byte(id), // send the summary id
	}); err != nil {
		deleteLocalTmpVideo(videoPath)
		s.storageRepository.DeleteFile(BUCKET_NAME, cloudRelativePath)
		s.summaryRepository.DeleteById(id)
		return nil, err
	}

	// Delete local tmp file
	if err = deleteLocalTmpVideo(videoPath); err != nil {
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

func NewSummaryService(summaryRepository ISummaryRepository, youtubeClient youtube.Client, taskPublisherRepository taskPublisher.ITaskPublisherRepository, storageRepository storage.IStorageRepository) ISummaryService {
	return &summaryService{
		summaryRepository:       summaryRepository,
		youtubeClient:           youtubeClient,
		taskPublisherRepository: taskPublisherRepository,
		storageRepository:       storageRepository,
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

func deleteLocalTmpVideo(videoPath string) error {
	if err := os.Remove(videoPath); err != nil {
		return err
	}
	return nil
}
