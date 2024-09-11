package services

import (
	"fmt"
	"log"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/infrastructures"
	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
)

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
		events, err := infrastructures.GoogleCalendarService.Events.List(room.CalendarId).
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

func GetTodayEvent() (eventList []model.Calendar, err error) {
	teacherMailAddress, err := repositories.GetTeacherMailAddress()
	if err != nil {
		return nil, fmt.Errorf("Fail to get today event:", err)
	}

	todayStart := time.Now().Truncate(24 * time.Hour).Format(time.RFC3339)
	todayEnd := time.Now().Truncate(24 * time.Hour).Add(23*time.Hour + 59*time.Minute + 59*time.Second).Format(time.RFC3339)

	eventIDs := make(map[string]bool)

	for _, teacherMail := range teacherMailAddress {
		events, err := infrastructures.GoogleCalendarService.Events.List(teacherMail).
			ShowDeleted(false).
			SingleEvents(true).
			TimeMin(todayStart).
			TimeMax(todayEnd).
			OrderBy("startTime").
			Do()
		if err != nil {
			log.Printf("Unable to retrieve events for teacher %s: %v", teacherMail, err)
			continue
		}

		for _, item := range events.Items {
			attendeeMail := []string{}
			for _, attendee := range item.Attendees {
				attendeeMail = append(attendeeMail, attendee.Email)
			}

			if len(attendeeMail) >= 7 {
				if _, exists := eventIDs[item.Id]; !exists {
					eventList = append(eventList, model.Calendar{
						RoomName:     "",
						Summary:      item.Summary,
						StartDate:    item.Start.DateTime,
						EndDate:      item.End.DateTime,
						AttendeeMail: attendeeMail,
					})
					eventIDs[item.Id] = true
				}
			}
		}
	}
	return eventList, nil
}
