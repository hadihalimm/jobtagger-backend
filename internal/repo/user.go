package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepo struct {
	db *config.Database
}

func NewUserRepo(db *config.Database) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Save(ctx context.Context, user *model.User) (*model.User, error) {
	query := `INSERT INTO users (full_name, email) VALUES ($1, $2) RETURNING id, full_name, email, created_at, updated_at`
	var savedUser model.User
	err := r.db.Pgx.QueryRow(ctx, query,
		user.FullName, user.Email).Scan(
		&savedUser.ID, &savedUser.FullName, &savedUser.Email, &savedUser.CreatedAt, &savedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &savedUser, nil
}

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `SELECT id, full_name, email FROM users WHERE id=$1`

	var user model.User
	err := r.db.Pgx.QueryRow(ctx, query, id).Scan(&user.ID, &user.FullName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, full_name, email FROM users WHERE email=$1`

	var user model.User
	err := r.db.Pgx.QueryRow(ctx, query, email).Scan(&user.ID, &user.FullName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
