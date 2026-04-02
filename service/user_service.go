package service

import (
	"context"
	"errors"
	"smp/models"
	"smp/repository"
	"smp/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	Repo  *repository.UserRepo
	Redis *redis.Client
}

func NewUserService(repo *repository.UserRepo, redis *redis.Client) *UserService {
	return &UserService{
		Repo:  repo,
		Redis: redis,
	}
}

func (s *UserService) OnboardUsers(ctx context.Context, user models.User) (string, error) {

	if user.Email == "" {
		return "", errors.New("email required")
	}

	if user.Password == "" || len(user.Password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	if user.Role != "teacher" && user.Role != "student" {
		return "", errors.New("please select your role")
	}

	// check if already verified
	verified, err := s.Repo.IsUserVerified(ctx, user.Email)
	if err == nil && verified {
		return "", errors.New("user already verified")
	}

	// OTP limit check
	if err := utils.CheckOTPLimit(ctx, s.Redis, user.Email); err != nil {
		return "", err
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	user.IsVerified = false
	user.CreatedAt = time.Now()

	userID, err := s.Repo.OnboardUsers(ctx, user)
	if err != nil {
		return "", err
	}

	otp := utils.GenerateOTP()

	err = utils.StoreOTP(ctx, s.Redis, user.Email, otp)
	if err != nil {
		return "", err
	}

	err = utils.SendOTPEmail(user.Email, otp)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *UserService) VerifyOTP(ctx context.Context, email string, otp string) error {

	if email == "" || otp == "" {
		return errors.New("email and otp required")
	}

	// get OTP from redis
	storedOTP, err := utils.GetOTP(ctx, s.Redis, email)
	if err != nil {
		return errors.New("otp expired")
	}

	// compare OTP
	if storedOTP != otp {
		return errors.New("invalid otp")
	}

	// mark user verified in DB
	err = s.Repo.VerifyUser(ctx, email)
	if err != nil {
		return err
	}

	// delete OTP after verification
	key := "otp:" + email
	s.Redis.Del(ctx, key)

	return nil
}

func (s *UserService) ResendOTP(ctx context.Context, email string) error {

	verified, err := s.Repo.IsUserVerified(ctx, email)
	if err != nil {
		return err
	}

	if verified {
		return errors.New("user already verified")
	}

	if err := utils.CheckOTPLimit(ctx, s.Redis, email); err != nil {
		return err
	}

	otp := utils.GenerateOTP()

	err = utils.StoreOTP(ctx, s.Redis, email, otp)
	if err != nil {
		return err
	}

	return utils.SendOTPEmail(email, otp)
}

func (s *UserService) Login(ctx context.Context, email string, password string) (string, error) {

	if email == "" || password == "" {
		return "", errors.New("email and password required")
	}

	user, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// check email verified
	if !user.IsVerified {
		return "", errors.New("email not verified")
	}

	// compare password
	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// generate jwt
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
