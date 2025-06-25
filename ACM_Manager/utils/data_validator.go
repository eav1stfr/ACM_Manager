package utils

import (
	"acmmanager/internal/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateMemberPost(newMembers []models.Member) error {
	for _, member := range newMembers {
		err := validate.Struct(member)
		if err != nil {
			return InvalidRequestPayloadError
		}
	}
	return nil
}
