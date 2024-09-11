package model

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
