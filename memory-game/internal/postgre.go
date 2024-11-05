package t_bot

import (
	"github.com/go-pg/pg"
)

type PostgreConfig struct {
	User     string
	Password string
	Port     string
	Host     string
}

type postgreStore struct {
	db *pg.DB
}

func PostgreUser(config PostgreConfig) UserInfo {
	db := pg.Connect(&pg.Options{
		Addr:     config.Host + ":" + config.Port,
		User:     config.User,
		Password: config.Password,
		Database: "postgres",
	})

	return &postgreStore{db: db}
}

func (p postgreStore) GetUser(id int) (*User, error) {
	user := &User{ID: id}
	err := p.db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p postgreStore) GetByExternalID(externalID string) (*User, error) {
	user := &User{}
	err := p.db.Model(&User{}).Where("external_id = ?", externalID).Select(user)
	return user, err
}

func (p postgreStore) GetAllUser(filter UserFilter) ([]*User, error) {
	var users []*User
	q := p.db.Model(&users).Order("score DESC")
	if filter.Limit > 0 {
		q.Limit(filter.Limit)
	}

	err := q.Select()

	if err != nil {
		return nil, err
	}
	return users, nil
}
func (p postgreStore) UpdateUser(id int, user *User) (*User, error) {
	user.ID = id

	err := p.db.Update(user)
	return user, err
}
func (p postgreStore) DeleteUser(id int) error {
	user := &User{ID: id}
	err := p.db.Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func (p postgreStore) CreateUser(user *User) (*User, error) {
	re := p.db.Insert(user)
	return user, re
}
