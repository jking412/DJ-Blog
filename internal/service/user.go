package service

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}
