package usecase

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"spectator.main/domain"
)

type configUsecase struct {
	configRepo     domain.ConfigRepository
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewConfigUsecase(c domain.ConfigRepository, u domain.UserRepository, to time.Duration) domain.ConfigUsecase {
	return &configUsecase{
		configRepo:     c,
		userRepo:       u,
		contextTimeout: to,
	}
}

func (c *configUsecase) InsertOne(ctx context.Context, config *domain.ConfigDetails) (*domain.ConfigDetails, error) {

	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	_, err := c.userRepo.FindOne(ctx, config.UserID.Hex())
	if err != nil {
		return nil, errors.New("user not found")
	}

	config.ID = primitive.NewObjectID()
	config.CreatedAt = time.Now()
	config.UpdatedAt = time.Now()

	res, err := c.configRepo.InsertOne(ctx, config)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *configUsecase) GetByUserID(ctx context.Context, userID string) (*domain.ConfigDetails, error) {

	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	res, err := c.configRepo.GetByUserID(ctx, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *configUsecase) AddSiteConfig(ctx context.Context, site_config *domain.SiteConfig, id string) error {

	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.configRepo.AddSiteConfig(ctx, site_config, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *configUsecase) RemoveSiteConfig(ctx context.Context, site_url string, id string) error {

	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.configRepo.RemoveSiteConfig(ctx, site_url, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *configUsecase) UpdateSiteConfig(ctx context.Context, site_config *domain.SiteConfig, id string) error {

	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.configRepo.UpdateSiteConfig(ctx, site_config, id)
	if err != nil {
		return err
	}

	return nil
}
