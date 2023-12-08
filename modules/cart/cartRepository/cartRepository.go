package cartRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	CartRepositoryService interface{}

	cartRepository struct {
		db *mongo.Client
	}
)

func NewCartRepository(db *mongo.Client) CartRepositoryService {
	return &cartRepository{db}
}

func (r *cartRepository) cartDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("cart_db")
}
