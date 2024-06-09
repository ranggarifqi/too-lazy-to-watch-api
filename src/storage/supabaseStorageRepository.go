package storage

import (
	"io"

	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase-community/supabase-go"
)

type supabaseStorageRepository struct {
	client *supabase.Client
}

func (s *supabaseStorageRepository) DeleteFile(bucketId string, relativePath string) error {
	_, err := s.client.Storage.RemoveFile(bucketId, []string{relativePath})

	if err != nil {
		return err
	}
	return nil
}

func (s *supabaseStorageRepository) Upload(bucketId string, relativePath string, data io.Reader, fileOptions FileOptions) (string, error) {
	upsert := true
	cacheControl := "3600"

	_, err := s.client.Storage.UploadFile(bucketId, relativePath, data, storage_go.FileOptions{
		ContentType:  &fileOptions.ContentType,
		Upsert:       &upsert,
		CacheControl: &cacheControl,
	})
	if err != nil {
		return "", err
	}

	result := s.client.Storage.GetPublicUrl(bucketId, relativePath)

	return result.SignedURL, nil
}

func NewSupabaseStorageRepository(client *supabase.Client) IStorageRepository {
	return &supabaseStorageRepository{
		client: client,
	}
}
