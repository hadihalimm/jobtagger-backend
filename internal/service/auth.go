package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hadihalimm/jobtagger-backend/internal/model"
	"github.com/hadihalimm/jobtagger-backend/internal/repo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AuthService interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	AuthCallback(w http.ResponseWriter, r *http.Request) (*model.User, error)
	SignOut(w http.ResponseWriter, r *http.Request) error
	GenerateAccessToken(id uuid.UUID) (string, error)
	GenerateRefreshToken(r *http.Request, id uuid.UUID) (string, error)
	ValidateRefreshToken(r *http.Request, incomingToken string) (*model.RefreshToken, error)
	RevokeRefreshToken(r *http.Request, refreshToken string) error
}

type authService struct {
	userRepo         repo.UserRepo
	refreshTokenRepo repo.RefreshTokenRepo
}

func NewAuthService(userRepo repo.UserRepo, refreshTokenRepo repo.RefreshTokenRepo) AuthService {
	return &authService{userRepo: userRepo, refreshTokenRepo: refreshTokenRepo}
}

func InitAuth() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), os.Getenv("GOOGLE_CALLBACK_URL")),
	)
}

func (s *authService) SignIn(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (s *authService) AuthCallback(w http.ResponseWriter, r *http.Request) (*model.User, error) {
	authUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return nil, err
	}
	var user *model.User
	user, err = s.userRepo.FindByEmail(r.Context(), authUser.Email)
	if err != nil {
		newUser := &model.User{FullName: authUser.Name, Email: authUser.Email}
		user, err = s.userRepo.Save(r.Context(), newUser)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *authService) SignOut(w http.ResponseWriter, r *http.Request) error {
	err := gothic.Logout(w, r)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) GenerateAccessToken(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *authService) GenerateRefreshToken(r *http.Request, id uuid.UUID) (string, error) {
	newRefreshToken := &model.RefreshToken{UserId: id, ExpiresAt: time.Now().Add(time.Hour * 24 * 14)}
	refreshToken, err := s.refreshTokenRepo.Save(r.Context(), newRefreshToken)
	if err != nil {
		return "", err
	}
	return refreshToken.Token.String(), nil
}

func (s *authService) ValidateRefreshToken(r *http.Request, incomingToken string) (*model.RefreshToken, error) {
	parsedToken, err := uuid.Parse(incomingToken)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.refreshTokenRepo.FindByToken(r.Context(), parsedToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		return nil, fmt.Errorf("refresh token expired")
	}

	return refreshToken, nil
}

func (s *authService) RevokeRefreshToken(r *http.Request, refreshToken string) error {
	parsedToken, err := uuid.Parse(refreshToken)
	if err != nil {
		return err
	}
	return s.refreshTokenRepo.Delete(r.Context(), parsedToken)
}
