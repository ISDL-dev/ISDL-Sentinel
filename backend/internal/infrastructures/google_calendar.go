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
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var GoogleCalendarConfig *oauth2.Config
var GoogleCalendarService *calendar.Service

func GoogleCalendarCallback(ctx *gin.Context) {
	authCode := ctx.Query("code")
	if authCode == "" {
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: "Code not found",
		})
		return
	}

	tok, err := GoogleCalendarConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	TokenChan <- tok
}

func InitializeGoogleCalendarClient() {
	log.Println("initialize Google Calendar client")
	googleCredentialsPath := os.Getenv("GOOGLE_CREDENTIALS_PATH")

	ctx := context.Background()
	b, err := os.ReadFile(fmt.Sprintf("%s/google_calendar_credentials.json", googleCredentialsPath))
	if err != nil {
		log.Fatalf("Failed to read the client secret file: %w", err)
	}

	calendarReadonlyScope := "https://www.googleapis.com/auth/calendar.readonly"
	GoogleCalendarConfig, err = google.ConfigFromJSON(b, calendarReadonlyScope)
	if err != nil {
		log.Fatalf("Failed to parse the client secret file: %w", err)
	}

	tokFile := fmt.Sprintf("%s/google_calendar_token.json", googleCredentialsPath)
	client := GetClient(GoogleCalendarConfig, tokFile)
	if err != nil {
		log.Fatalf("Failed to get client: %v", err)
	}

	GoogleCalendarService, err = calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Failed to initialize Google Calendar client: %w", err)
	}
}
