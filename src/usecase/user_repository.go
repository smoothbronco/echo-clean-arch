package usecase

import "github.com/smoothbronco/echo-clean-arch/src/domain"

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}