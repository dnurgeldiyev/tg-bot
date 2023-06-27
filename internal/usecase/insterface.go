package usecase

import "dovran/tg-bot/internal/entity"

type UseCase interface {
	InsertUser(user entity.UserData) error
	BulkInsert(userList []entity.UserData) (int, error)
	SearchByName(s string) (entity.UserData, error)
	GetByLastNDate(n int) ([]entity.UserData, error)
}
