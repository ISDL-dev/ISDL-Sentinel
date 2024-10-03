package services

import (
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
    "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/schema"
)

func PostSignUpService(user schema.PostUserSignUpRequest) error {
	return repositories.SignUpRepository(user)
}