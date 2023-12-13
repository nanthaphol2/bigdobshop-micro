package migration

import (
	"context"
	"log"

	"github.com/nanthaphol2/bigdobshop-micro/config"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/database"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func cartDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("cart_db")
}

func CartMigrate(pctx context.Context, cfg *config.Config) {
	db := cartDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("carts")
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})
	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	documents := func() []any {
		roles := []*cart.Cart{
			{
				Title:       "Diamond Sword",
				Price:       1000,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Iron Sword",
				Price:       500,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Wooden Sword",
				Price:       100,
				ImageUrl:    "https://i.imgur.com/1Y8tQZM.png",
				UsageStatus: true,
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
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
	log.Println("Migrate cart completed: ", results)
}
