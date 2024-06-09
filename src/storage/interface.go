package storage

import "io"

type FileOptions struct {
	ContentType string
}

type IStorageRepository interface {
	Upload(bucketId string, relativePath string, data io.Reader, fileOptions FileOptions) (string, error)
}
