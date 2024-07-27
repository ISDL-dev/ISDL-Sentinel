package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

const webhookURL = "https://prod-22.japaneast.logic.azure.com:443/workflows/baaff47922d44cf1b94d32c1392f965f/triggers/manual/paths/invoke?api-version=2016-06-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=k7eFfJ8C0-3XLKbmGW6_mi31GGTHE2iXTktmNqmySyo"

type Content struct {
	Schema  string `json:"$schema"`
	Type    string `json:"type"`
	Version string `json:"version"`
	Body    []Body `json:"body"`
}

type Body struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Attachment struct {
	ContentType string  `json:"contentType"`
	ContentURL  *string `json:"contentUrl"`
	Content     Content `json:"content"`
}

type Message struct {
	Type        string       `json:"type"`
	Attachments []Attachment `json:"attachments"`
}

func PostWebHookTest(user schema.Status) {
	message := Message{
		Type: "message",
		Attachments: []Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentURL:  nil,
				Content: Content{
					Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
					Type:    "AdaptiveCard",
					Version: "1.4",
					Body: []Body{
						{
							Type: "TextBlock",
							Text: fmt.Sprintf("%d が %s しました", user.UserId, user.Status),
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

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
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
