package usecase

import (
	"context"
	"time"

	"spectator.main/domain"
	tokenutil "spectator.main/internals/util"
)

type authUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		userRepo:       userRepository,
		contextTimeout: timeout,
	}
}

func (au *authUsecase) CreateUser(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	_, err := au.userRepo.InsertOne(ctx, user)
	return err
}

func (au *authUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.userRepo.FindByEmail(ctx, email)
}

func (au *authUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (au *authUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
