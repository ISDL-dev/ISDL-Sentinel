package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	"google.golang.org/api/drive/v3"
)

func UploadAvatarFile(avatarFile *multipart.FileHeader) (string, error) {
	file, err := avatarFile.Open()
	if err != nil {
		return "", fmt.Errorf("Failed to open the file: %w", err)
	}
	defer file.Close()

	fileMetadata := &drive.File{
		Name:    avatarFile.Filename,
		Parents: []string{os.Getenv("GOOGLE_DRIVE_FOLDER_ID")},
	}

	driveFile, err := infrastructures.GoogleDriveService.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("Failed to upload the file to Google Drive: %w", err)
	}

	avatarImgPath := fmt.Sprintf("https://drive.google.com/thumbnail?id=%s&sz=w1000", driveFile.Id)

	return avatarImgPath, nil
}

func DeleteAvatarFile(avatarImgPath string) error {
	const idParam = "id="
	idIndex := strings.Index(avatarImgPath, idParam)
	if idIndex == -1 {
		return fmt.Errorf("File ID not found in avatarImgPath")
	}
	idStart := idIndex + len(idParam)
	idEnd := strings.Index(avatarImgPath[idStart:], "&")
	if idEnd == -1 {
		idEnd = len(avatarImgPath)
	} else {
		idEnd += idStart
	}
	fileID := avatarImgPath[idStart:idEnd]

	err := infrastructures.GoogleDriveService.Files.Delete(fileID).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete the file from Google Drive: %w", err)
	}

	return nil
}
