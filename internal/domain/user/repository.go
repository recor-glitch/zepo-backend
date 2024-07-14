package user

type Repository interface {
	GetByID(id string) (*User, error)
	Create(name *User) (*User, error)
}