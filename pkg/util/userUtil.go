package util

import (
	"intern-bcc/model"

	"github.com/google/uuid"
)

func RemoveUser(users []model.UserFilter) []model.UserFilter {
	dUser := map[uuid.UUID]bool{}
	result := []model.UserFilter{}

	for _, user := range users {
		if dUser[user.ID] {
			continue
		}
		dUser[user.ID] = true
		result = append(result, user)
	}
	return result
}
