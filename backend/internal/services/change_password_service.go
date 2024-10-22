package services

import (
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PutChangePasswordService(user schema.PutChangePasswordRequest, userID int) error {
    return repositories.ChangePasswordRepository(user, userID)
}