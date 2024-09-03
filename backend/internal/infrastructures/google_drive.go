package infrastructures

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var googleDriveConfig *oauth2.Config
var googleDriveService *drive.Service

func GoogleDriveCallback(ctx *gin.Context) {
	authCode := ctx.Query("code")
	if authCode == "" {
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: "Code not found",
		})
		return
	}

	tok, err := googleDriveConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	TokenChan <- tok
}

func InitializeGoogleDriveClient() {
	log.Println("initialize Google Drive client")
	ctx := context.Background()
	b, err := os.ReadFile("internal/infrastructures/credentials/google_drive_credentials.json")
	if err != nil {
		fmt.Errorf("Failed to read the client secret file: %w", err)
	}

	googleDriveConfig, err = google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		fmt.Errorf("Failed to parse the client secret file: %w", err)
	}

	tokFile := "internal/infrastructures/credentials/google_drive_token.json"
	client := GetClient(googleDriveConfig, tokFile)

	googleDriveService, err = drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Errorf("Failed to initialize Google Drive client: %w", err)
	}
}

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

	driveFile, err := googleDriveService.Files.Create(fileMetadata).Media(file).Do()
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

	err := googleDriveService.Files.Delete(fileID).Do()
	if err != nil {
		return fmt.Errorf("Failed to delete the file from Google Drive: %w", err)
	}

	return nil
}
