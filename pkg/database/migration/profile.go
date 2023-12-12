package migration

import (
	"context"
	"log"

	"github.com/nanthaphol2/bigdobshop-micro/config"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/database"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func profileDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("profile_db")
}

func ProfileMigrate(pctx context.Context, cfg *config.Config) {
	db := profileDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("profile_transactions")

	// indexs
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"profile_id", 1}}},
	})
	log.Println(indexs)

	col = db.Collection("profiles")

	// indexs
	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"email", 1}}},
	})
	log.Println(indexs)

	documents := func() []any {
		roles := []*profile.Profile{
			{
				Email: "u001@bgshop.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "u001",
				ProfileRoles: []profile.ProfileRole{
					{
						RoleTitle: "user",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "a001@bgshop.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "a001",
				ProfileRoles: []profile.ProfileRole{
					{
						RoleTitle: "user",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate auth completed: ", results)

	profileTransactions := make([]any, 0)
	for _, p := range results.InsertedIDs {
		profileTransactions = append(profileTransactions, &profile.ProfileTransaction{
			ProfileId: p.(primitive.ObjectID).Hex(),
			Amount:    1000,
			CreatedAt: utils.LocalTime(),
		})
	}

	col = db.Collection("profile_transactions")
	results, err = col.InsertMany(pctx, profileTransactions, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate profile_transactions completed: ", results)

	col = db.Collection("profile_transactions_queue")
	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate profile_transactions_queue completed: ", result)
}
