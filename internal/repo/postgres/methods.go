package postgres

import "dovran/tg-bot/internal/entity"

func (p *Postgres) Insert(user entity.UserData) error {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) BulkInsert(userList []entity.UserData) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) SearchByName(s string) (entity.UserData, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) GetByLastNDate(n int) ([]entity.UserData, error) {
	//TODO implement me
	panic("implement me")
}
