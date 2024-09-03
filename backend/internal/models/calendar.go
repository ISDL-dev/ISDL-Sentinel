package model

type Calendar struct {
	RoomName     string
	Summary      string
	StartDate    string
	EndDate      string
	AttendeeMail []string
}

type CalendarRoomId struct {
	RoomName   string
	CalendarId string
}

const (
	KC101_LARGE_CALENDAR_ID = "mikilab.doshisha.ac.jp_3739313235333736353437@resource.calendar.google.com"
	KC101_SMALL_CALENDAR_ID = "c_188c9tphie1akjm2hquasoipu060q@resource.calendar.google.com"
	KC103_CALENDAR_ID       = "mikilab.doshisha.ac.jp_33353234353936362d333132@resource.calendar.google.com"
	KC111_CALENDAR_ID       = "mikilab.doshisha.ac.jp_3235333239333534343633@resource.calendar.google.com"
	KC116_CALENDAR_ID       = "mikilab.doshisha.ac.jp_38363338373137302d343939@resource.calendar.google.com"
	KC119_CALENDAR_ID       = "mikilab.doshisha.ac.jp_38363935323038372d393739@resource.calendar.google.com"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}

func ToInterfaceSlice(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
