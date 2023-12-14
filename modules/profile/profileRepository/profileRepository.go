package profileRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/nanthaphol2/bigdobshop-micro/modules/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ProfileRepositoryService interface {
		IsUniqueProfile(pctx context.Context, email, username string) bool
		InsertOneProfile(pctx context.Context, req *profile.Profile) (primitive.ObjectID, error)
	}

	profileRepository struct {
		db *mongo.Client
	}
)

func NewProfileRepository(db *mongo.Client) ProfileRepositoryService {
	return &profileRepository{db: db}
}

func (r *profileRepository) profileDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("profile_db")
}

func (r *profileRepository) IsUniqueProfile(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.profileDbConn(ctx)
	col := db.Collection("profiles")

	player := new(profile.Profile)
	if err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{"username": username},
			{"email": email},
		}},
	).Decode(player); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}
	return false
}

func (r *profileRepository) InsertOneProfile(pctx context.Context, req *profile.Profile) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.profileDbConn(ctx)
	col := db.Collection("profiles")

	playerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOneProfile: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one profile failed")
	}

	return playerId.InsertedID.(primitive.ObjectID), nil
}
