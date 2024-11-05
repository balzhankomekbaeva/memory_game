package t_bot

type UserInfo interface {
	GetUser(id int) (*User, error)
	GetByExternalID(externalID string) (*User, error)
	UpdateUser(id int, user *User) (*User, error)
	DeleteUser(id int) error
	GetAllUser(filter UserFilter) ([]*User, error)
	CreateUser(user *User) (*User, error)
}
