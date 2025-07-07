package utils

import (
	"acmmanager/internal/models"
	"github.com/go-playground/validator/v10"
	"log"
)

var validate = validator.New()

func ValidatePostMember(newMember models.Member) error {
	err := validate.Struct(newMember)
	if err != nil {
		return InvalidRequestPayloadError
	}
	return nil
}
func ValidateMemberPost(newMembers []models.Member) error {
	for _, member := range newMembers {
		err := validate.Struct(member)
		if err != nil {
			return InvalidRequestPayloadError
		}
	}
	return nil
}

func ValidateTaskPost(newTask models.Task) error {
	err := validate.Struct(newTask)
	if err != nil {
		return InvalidRequestPayloadError
	}
	return nil
}

func ValidateMeetingPost(newMeeting models.Meeting) error {
	err := validate.Struct(newMeeting)
	if err != nil {
		log.Println("ERR IS", err)
		return InvalidRequestPayloadError
	}
	return nil
}
