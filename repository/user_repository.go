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
	INSERT INTO users (email,password,role,is_verified,created_at)
	VALUES ($1,$2,$3,false,$4)
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
