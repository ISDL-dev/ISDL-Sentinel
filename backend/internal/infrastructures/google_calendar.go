package infrastructures

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var googleCalendarConfig *oauth2.Config
var googleCalendarService *calendar.Service

func GoogleCalendarCallback(ctx *gin.Context) {
	authCode := ctx.Query("code")
	if authCode == "" {
		ctx.JSON(http.StatusBadRequest, schema.Error{
			Code:    http.StatusBadRequest,
			Message: "Code not found",
		})
		return
	}

	tok, err := googleCalendarConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	TokenChan <- tok
}

func InitializeGoogleCalendarClient() {
	log.Println("initialize Google Calendar client")
	ctx := context.Background()
	b, err := os.ReadFile("internal/infrastructures/credentials/google_calendar_credentials.json")
	if err != nil {
		fmt.Errorf("Failed to read the client secret file: %w", err)
	}

	calendarReadonlyScope := "https://www.googleapis.com/auth/calendar.readonly"
	googleCalendarConfig, err = google.ConfigFromJSON(b, calendarReadonlyScope)
	if err != nil {
		fmt.Errorf("Failed to parse the client secret file: %w", err)
	}

	tokFile := "internal/infrastructures/credentials/google_calendar_token.json"
	client := GetClient(googleCalendarConfig, tokFile)

	googleCalendarService, err = calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Errorf("Failed to initialize Google Calendar client: %w", err)
	}
}

func GetCalendarList() []model.Calendar {
	var eventList []model.Calendar
	calendarIDs := []model.CalendarRoomId{
		model.CalendarRoomId{RoomName: "KC101-large", CalendarId: model.KC101_LARGE_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC101-small", CalendarId: model.KC101_SMALL_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC103", CalendarId: model.KC103_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC111", CalendarId: model.KC111_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC116", CalendarId: model.KC116_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC119", CalendarId: model.KC119_CALENDAR_ID},
	}

	// 現在の時刻を取得
	currentTime := time.Now().Format(time.RFC3339)

	// 各カレンダーからイベントを取得
	for _, room := range calendarIDs {
		events, err := googleCalendarService.Events.List(room.CalendarId).
			ShowDeleted(false).
			SingleEvents(true).
			TimeMin(currentTime).
			MaxResults(2).
			OrderBy("startTime").
			Do()
		if err != nil {
			log.Printf("Unable to retrieve events for calendar %s: %v", room.RoomName, err)
			continue
		}

		if len(events.Items) == 0 {
			fmt.Printf("No upcoming events found for calendar %s.\n", room.RoomName)
		} else {
			fmt.Printf("Upcoming events for calendar %s:\n", room.RoomName)
			for _, item := range events.Items {
				attendeeMail := []string{}
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				for _, attendee := range item.Attendees {
					if attendee.Email == room.CalendarId {
						continue
					}
					attendeeMail = append(attendeeMail, attendee.Email)
				}
				eventList = append(eventList, model.Calendar{
					RoomName:     room.RoomName,
					Summary:      item.Summary,
					StartDate:    item.Start.DateTime,
					EndDate:      item.End.DateTime,
					AttendeeMail: attendeeMail,
				})
			}
		}
	}
	return eventList
}
