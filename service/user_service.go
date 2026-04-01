package service

import (
	"context"
	"smp/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{Repo: repo}
}


func(s *UserService) OnboardUsers (ctx context.Context, )