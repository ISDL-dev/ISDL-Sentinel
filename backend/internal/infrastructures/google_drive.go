package infrastructures

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var GoogleDriveConfig *oauth2.Config
var GoogleDriveService *drive.Service

func GoogleDriveCallback(ctx *gin.Context) {
	authCode := ctx.Query("code")
	if authCode == "" {
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: "Code not found",
		})
		return
	}

	tok, err := GoogleDriveConfig.Exchange(context.TODO(), authCode)
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

	GoogleDriveConfig, err = google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		fmt.Errorf("Failed to parse the client secret file: %w", err)
	}

	tokFile := "internal/infrastructures/credentials/google_drive_token.json"
	client := GetClient(GoogleDriveConfig, tokFile)

	GoogleDriveService, err = drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Errorf("Failed to initialize Google Drive client: %w", err)
	}
}
