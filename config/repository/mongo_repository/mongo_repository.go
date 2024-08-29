package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"spectator.main/domain"
	"spectator.main/internals/mongo"
)

type mongoRepository struct {
	DB         mongo.Database
	Collection mongo.Collection
}

const (
	timeFormat     = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
	collectionName = "config"
)

func NewMongoRepository(DB mongo.Database) domain.ConfigRepository {
	return &mongoRepository{DB, DB.Collection(collectionName)}
}

func (m *mongoRepository) InsertOne(ctx context.Context, config *domain.ConfigDetails) (*domain.ConfigDetails, error) {
	var (
		err error
	)

	_, err = m.Collection.InsertOne(ctx, config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (m *mongoRepository) GetByUserID(ctx context.Context, userID string) (*domain.ConfigDetails, error) {
	var (
		config domain.ConfigDetails
		err    error
	)

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return &config, err
	}

	err = m.Collection.FindOne(ctx, bson.M{"user_id": idHex}).Decode(&config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}

func (m *mongoRepository) AddSiteConfig(ctx context.Context, site_config *domain.SiteConfig, id string) error {
	var (
		err error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": idHex}

	update := bson.M{
		"$push": bson.M{
			"site_configs": site_config,
		},
	}

	_, err = m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err

	}

	return nil
}

func (m *mongoRepository) RemoveSiteConfig(ctx context.Context, site_url string, id string) error {

	var (
		err error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	fmt.Println(site_url)
	filter := bson.M{"_id": idHex}

	update := bson.M{
		"$pull": bson.M{
			"site_configs": bson.M{
				"site_url": site_url,
			},
		},
	}

	result, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("no site config found with the given site url")
	}

	return nil
}

func (m *mongoRepository) UpdateSiteConfig(ctx context.Context, site_config *domain.SiteConfig, id string) error {

	var (
		err error
	)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	fmt.Println(site_config.SiteUrl)
	filter := bson.M{"_id": idHex, "site_configs.site_url": site_config.SiteUrl}

	update := bson.M{
		"$set": bson.M{
			"site_configs.$": site_config,
		},
	}

	result, err := m.Collection.UpdateOne(ctx, filter, update)

	if result.ModifiedCount == 0 {
		return errors.New("no site config found with the given site url")
	}

	if err != nil {
		return err
	}

	return nil

}
