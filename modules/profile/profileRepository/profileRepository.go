package profileRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ProfileRepositoryService interface{}

	profileRepository struct {
		db *mongo.Client
	}
)

func NewProfileRepository(db *mongo.Client) ProfileRepositoryService {
	return &profileRepository{db}
}

func (r *profileRepository) profileDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("profile_db")
}
