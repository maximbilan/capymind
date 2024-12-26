package filestorage

import "time"

type FileStorage interface {
	Upload(filePath string, expirationDate time.Time) *string
}
