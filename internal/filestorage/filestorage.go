package filestorage

import "time"

type FileStorage interface {
	Upload(title string, filePath string, expirationDate time.Time) *string
}
