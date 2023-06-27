package usecase

import (
	"dovran/tg-bot/internal/entity"
	"dovran/tg-bot/internal/repo"
)

type Ucase struct {
	u repo.Repo
}

func NewUcase(u repo.Repo) UseCase {
	return Ucase{u: u}
}

func (u Ucase) InsertUser(user entity.UserData) error {
	//TODO implement me
	panic("implement me")
}

func (u Ucase) BulkInsert(userList []entity.UserData) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u Ucase) SearchByName(s string) (entity.UserData, error) {
	//TODO implement me
	panic("implement me")
}

func (u Ucase) GetByLastNDate(n int) ([]entity.UserData, error) {
	//TODO implement me
	panic("implement me")
}
