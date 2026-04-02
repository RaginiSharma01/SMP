package repository

import (
	"context"
	"errors"
	"smp/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		DB: pool,
	}
}

func (r *UserRepo) OnboardUsers(ctx context.Context, user models.User) (string, error) {

	query := `
	INSERT INTO users (email, password, role, is_verified, created_at)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		user.Email,
		user.Password,
		user.Role,
		user.IsVerified,
		user.CreatedAt,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepo) VerifyUser(ctx context.Context, email string) error {

	query := `
	UPDATE users
	SET is_verified = true
	WHERE email = $1
	`

	_, err := r.DB.Exec(ctx, query, email)
	return err
}

func (r *UserRepo) IsUserVerified(ctx context.Context, email string) (bool, error) {

	query := `
	SELECT is_verified
	FROM users
	WHERE email = $1
	`

	var verified bool

	err := r.DB.QueryRow(ctx, query, email).Scan(&verified)

	if err == pgx.ErrNoRows {
		return false, errors.New("user not found")
	}

	if err != nil {
		return false, err
	}

	return verified, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	query := `
	SELECT id, email, password, role, is_verified
	FROM users
	WHERE email = $1
	`

	var user models.User

	err := r.DB.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsVerified,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
