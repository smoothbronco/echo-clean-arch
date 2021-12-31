package usecase

import "github.com/smoothbronco/echo-clean-arch/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) GetOneInfo(id int) domain.User {
	return interactor.UserRepository.SelectById(id)
}

func (interactor *UserInteractor) Delete(id int) {
	interactor.UserRepository.Delete(id)
}
