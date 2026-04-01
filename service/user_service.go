package service

import (
	"context"
	"errors"
	"smp/models"
	"smp/repository"
	"smp/utils"
	"time"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) OnboardUsers(ctx context.Context, user models.User) (string, error) {

	if user.Email == "" {
		return "", errors.New("email required")
	}

	if user.Password == "" || len(user.Password) < 8 {
		return "", errors.New("password must be atleast 8 characters")
	}

	if user.Role != "teacher" && user.Role != "student" {
		return "", errors.New("Please select your role")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	user.IsVerified = false
	user.CreatedAt = time.Now()

	// save user
	userID, err := s.Repo.OnboardUsers(ctx, user)
	if err != nil {
		return "", err
	}

	// generate OTP
	otp := utils.GenerateOTP()

	// send email
	err = utils.SendOTPEmail(user.Email, otp)
	if err != nil {
		return "", err
	}

	return userID, nil
}
