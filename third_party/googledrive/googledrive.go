package googledrive

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleDrive struct{}

func credentialsPath() string {
	var path = "credentials.json"
	if os.Getenv("DEBUG_MODE") == "true" {
		path = "./../" + path
	} else {
		path = "./" + path
	}
	return path
}

func create(ctx context.Context) (*drive.Service, error) {
	var srv *drive.Service
	var err error

	if os.Getenv("CLOUD") == "true" {
		srv, err = drive.NewService(ctx)
	} else {
		srv, err = drive.NewService(ctx, option.WithCredentialsFile(credentialsPath()))
	}

	if err != nil {
		return nil, fmt.Errorf("unable to create Drive service: %w", err)
	}
	return srv, nil
}

func (googleDrive GoogleDrive) Upload(title string, filePath string, expirationDate time.Time) *string {
	ctx := context.Background()

	srv, err := create(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return nil
	}
	defer file.Close()

	driveFile := &drive.File{
		Name:    title,
		Parents: []string{"root"},
	}

	// Upload the file using the Files.Create method
	uploadedFile, err := srv.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		return nil
	}

	// Set the file to be publicly accessible
	_, err = srv.Permissions.Create(uploadedFile.Id, &drive.Permission{
		Type:           "anyone",
		Role:           "reader",
		ExpirationTime: expirationDate.Format(time.RFC3339),
	}).Do()
	if err != nil {
		log.Printf("Error setting file permissions: %v\n", err)
		return nil
	}

	// Generate a shareable link
	link := fmt.Sprintf("https://drive.google.com/file/d/%s/view", uploadedFile.Id)
	return &link
}
