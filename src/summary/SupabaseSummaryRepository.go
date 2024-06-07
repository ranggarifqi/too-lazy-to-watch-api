package summary

import (
	"encoding/json"
	"fmt"
	"os"

	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase-community/supabase-go"
)

type supabaseSummaryRepository struct {
	client *supabase.Client
}

const BUCKET_NAME = "video"

func (s *supabaseSummaryRepository) UploadVideo(tmpVideoPath string, uniqueId string) (string, error) {
	// Open the video file
	videoFile, err := os.Open(tmpVideoPath)
	if err != nil {
		return "", err
	}
	defer videoFile.Close()

	contentType := "video/mp4"
	upsert := true
	cacheControl := "3600"
	cloudRelativePath := fmt.Sprintf("%s.mp4", uniqueId)

	_, err = s.client.Storage.UploadFile(BUCKET_NAME, cloudRelativePath, videoFile, storage_go.FileOptions{
		ContentType:  &contentType,
		Upsert:       &upsert,
		CacheControl: &cacheControl,
	})
	if err != nil {
		return "", err
	}

	fmt.Printf("Uploaded")

	result := s.client.Storage.GetPublicUrl(BUCKET_NAME, cloudRelativePath)

	return result.SignedURL, nil
}

func (s *supabaseSummaryRepository) Create(payload CreateSummaryPayload) (*Summary, error) {
	res, _, err := s.client.From(TABLE_NAME).Insert(payload, false, "", "", "").Execute()
	if err != nil {
		return nil, err
	}

	summary := []Summary{}
	err = json.Unmarshal(res, &summary)
	if err != nil {
		return nil, err
	}

	return &summary[0], nil
}

func NewSupabaseSummaryRepository(client *supabase.Client) ISummaryRepository {
	return &supabaseSummaryRepository{
		client: client,
	}
}
