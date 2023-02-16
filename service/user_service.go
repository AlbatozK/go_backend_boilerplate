package service

import (
	"github.com/AlbatozK/go_backend_boilerplate/model"
	"github.com/AlbatozK/go_backend_boilerplate/repository"
)

type UserService struct {
	genRepo *repository.GenericRepository
}

func NewUserService() *UserService {
	return &UserService{
		genRepo: repository.NewGenericRepository(),
	}
}

func (us *UserService) GetUser(id int) (*model.User, error) {
	err := us.genRepo.BeginTx()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			us.genRepo.Rollback()
		} else {
			us.genRepo.Commit()
		}
	}()
	userRepo := us.genRepo.GetUserRepository()
	user, err := userRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
