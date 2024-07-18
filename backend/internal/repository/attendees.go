package repository

import (
	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/schema"
)

func GetAttendeesList() ([]schema.AttendeesListInner) {
	rows_attendee, err := db.Query("SELECT id, google_drive_id FROM images ORDER BY RAND() LIMIT ?", numImages)
}