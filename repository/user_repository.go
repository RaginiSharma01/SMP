package repository

import (
	"context"
	"smp/models"

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
