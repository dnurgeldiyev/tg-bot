package repo

import "dovran/tg-bot/internal/entity"

type Repo interface {
	Insert(user entity.UserData) error
	BulkInsert(userList []entity.UserData) (int, error)
	SearchByName(s string) (entity.UserData, error)
	GetByLastNDate(n int) ([]entity.UserData, error)
}
