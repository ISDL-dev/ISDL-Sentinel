package model

const (
	IN_ROOM = "In Room"
	OUT_ROOM = "Out of Room"
	OVERNIGHT = "Overnight"
)

func IsInRoom(status string) bool {
	return status != IN_ROOM
}