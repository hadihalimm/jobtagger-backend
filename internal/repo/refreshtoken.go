package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
)

type RefreshTokenRepo interface {
	Save(ctx context.Context, token *model.RefreshToken) (*model.RefreshToken, error)
	FindByToken(ctx context.Context, token uuid.UUID) (*model.RefreshToken, error)
	Delete(ctx context.Context, token uuid.UUID) error
}

type refreshTokenRepo struct {
	db *config.Database
}

func NewRefreshTokenRepo(db *config.Database) RefreshTokenRepo {
	return &refreshTokenRepo{db: db}
}

func (r *refreshTokenRepo) Save(ctx context.Context, token *model.RefreshToken) (*model.RefreshToken, error) {
	query := `INSERT INTO refresh_tokens (user_id, expires_at) VALUES ($1, $2)  RETURNING token, user_id, expires_at, created_at`
	var savedToken model.RefreshToken
	err := r.db.Pgx.QueryRow(ctx, query,
		token.UserId, token.ExpiresAt).Scan(
		&savedToken.Token, &savedToken.UserId, &savedToken.ExpiresAt, &savedToken.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &savedToken, nil
}

func (r *refreshTokenRepo) FindByToken(ctx context.Context, token uuid.UUID) (*model.RefreshToken, error) {
	query := `SELECT token, user_id, expires_at FROM refresh_tokens WHERE token=$1`
	var savedToken model.RefreshToken
	err := r.db.Pgx.QueryRow(ctx, query, token).Scan(
		&savedToken.Token, &savedToken.UserId, &savedToken.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &savedToken, nil
}

func (r *refreshTokenRepo) Delete(ctx context.Context, token uuid.UUID) error {
	query := `DELETE FROM refresh_tokens WHERE token=$1`
	_, err := r.db.Pgx.Exec(ctx, query, token)
	if err != nil {
		return err
	}
	return nil
}
