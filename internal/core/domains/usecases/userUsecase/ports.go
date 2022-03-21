package userUsecase

type UserService interface {
	Exec() error
}

type UserExternalService interface {
	Exec() error
}
