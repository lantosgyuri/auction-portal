package domain

type UserEvent interface {
	GetName() string
}

type CreateUserRequested struct {
	Name     string
	Password string
}

type DeleteUserRequest struct {
	Id   int
	Name string
}

type User struct {
	Id       int
	Name     string
	Password string
}

type UserEventRaw struct {
	EventType string
	UserId    int
	Name      string
	Password  string
	TimeStamp int
}

func (c CreateUserRequested) GetName() string {
	return c.Name
}

func (d DeleteUserRequest) GetName() string {
	return d.Name
}
