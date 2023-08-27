package utils

import (
	"fmt"
	"github.com/dailyplan-api/pkg/repository"
)

func GetCredentialsByRole(role string) ([]string, error) {
	var credentials []string

	switch role {
	case repository.AdminRoleName:
		credentials = []string{
			repository.HabitCreateCredential,
			repository.HabitUpdateCredential,
			repository.HabitDeleteCredential,
		}
	case repository.ModeratorRoleName:
		credentials = []string{
			repository.HabitCreateCredential,
			repository.HabitUpdateCredential,
		}
	case repository.UserRoleName:
		credentials = []string{
			repository.HabitCreateCredential,
		}
	default:
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
