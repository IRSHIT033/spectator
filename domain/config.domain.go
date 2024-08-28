package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConfigDetails struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	SiteConfig []SiteConfig       `bson:"site_configs" json:"site_configs"`
}

type SiteConfig struct {
	SiteUrl       string          `bson:"site_url" json:"site_url" validate:"required"`
	RegionDetails []RegionDetails `bson:"region_details" json:"region_details"`
}

type RegionDetails struct {
	Status       bool      `bson:"status" json:"status"`
	Region       string    `bson:"region" json:"region"`
	ResponseTime time.Time `bson:"response_time" json:"response_time"`
}

type RemoveConfigRequest struct {
	SiteUrl string `json:"site_url" validate:"required"`
}

type ConfigRepository interface {
	InsertOne(ctx context.Context, config *ConfigDetails) (*ConfigDetails, error)
	UpdateSiteConfig(ctx context.Context, site_config *SiteConfig, id string) error
	RemoveSiteConfig(ctx context.Context, site_url string, id string) error
	GetByUserID(ctx context.Context, userID string) (*ConfigDetails, error)
	AddSiteConfig(ctx context.Context, site_config *SiteConfig, id string) (*ConfigDetails, error)
}

type ConfigUsecase interface {
	InsertOne(ctx context.Context, config *ConfigDetails) (*ConfigDetails, error)
	UpdateSiteConfig(ctx context.Context, site_config *SiteConfig, id string) error
	RemoveSiteConfig(ctx context.Context, site_url string, id string) error
	GetByUserID(ctx context.Context, userID string) (*ConfigDetails, error)
	AddSiteConfig(ctx context.Context, site_config *SiteConfig, id string) (*ConfigDetails, error)
}
