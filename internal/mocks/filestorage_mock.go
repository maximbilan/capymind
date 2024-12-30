package mocks

import "time"

type ValidFileStorageMock struct{}

func (googleDrive ValidFileStorageMock) Upload(title string, filePath string, expirationDate time.Time) *string {
	string := "link"
	return &string
}

type InvalidFileStorageMock struct{}

func (googleDrive InvalidFileStorageMock) Upload(title string, filePath string, expirationDate time.Time) *string {
	return nil
}
