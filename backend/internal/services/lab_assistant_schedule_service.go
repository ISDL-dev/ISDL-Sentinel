package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func GetLabAssistantScheduleService(month string) (labAssistantSchedule []schema.GetLabAssistantSchedule200ResponseInner, err error) {
	labAssistantSchedule, err = repositories.GetLabAssistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant schedule: %v", err)
	}

	return labAssistantSchedule, nil
}

func PostLabAssistantScheduleService(month string, labAssistantScheduleRequest []schema.PostLabAssistantScheduleRequestInner) (labAssistantSchedule []schema.GetLabAssistantSchedule200ResponseInner, err error) {
	err = repositories.PostLabAssistantScheduleRepository(month, labAssistantScheduleRequest)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to post lab assistant schedule: %v", err)
	}

	labAssistantSchedule, err = repositories.GetLabAssistantScheduleRepository(month)
	if err != nil {
		return []schema.GetLabAssistantSchedule200ResponseInner{}, fmt.Errorf("failed to execute query to get lab assistant schedule: %v", err)
	}

	return labAssistantSchedule, nil
}

func NotificationLabAssistantScheduleWithTeams(shiftDate string, userName string) {
	event, err := GetTodayEvent()
	if err != nil {
		log.Println(err)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", shiftDate)
	if err != nil {
		log.Println("Date parsing error:", err)
		return
	}
	weekday := model.Weekdays[parsedDate.Weekday()]
	formattedDate := fmt.Sprintf("%d月%d日（%s）", parsedDate.Month(), parsedDate.Day(), weekday)

	var eventDetails string
	if len(event) == 0 {
		eventDetails = "ありません．"
	} else {
		eventDetails = "\n\n━━━━━━━━━━━━━━━━━\n\n"
		for _, e := range event {
			eventDetails += fmt.Sprintf("・%s\n\n", e.Summary)
		}
		eventDetails += "━━━━━━━━━━━━━━━━━"
	}

	message := model.Message{
		Type: "message",
		Attachments: []model.Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentURL:  nil,
				Content: model.Content{
					Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
					Type:    "AdaptiveCard",
					Version: "1.4",
					Body: []model.Body{
						{
							Type: "TextBlock",
							Text: fmt.Sprintf("おはようございます．\n\n本日，%sのLAは%sです．\n\n本日の行事は%s\n\n今日も一日頑張りましょう．", formattedDate, userName, eventDetails),
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", os.Getenv("WEBHOOKURL"), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received non-OK response:", resp.Status)
	} else {
		fmt.Println("Message sent successfully")
	}
}
