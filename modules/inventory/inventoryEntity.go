package inventory

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Inventory struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		ProfileId string             `json:"profile_id" bson:"profile_id"`
		CartId    string             `json:"cart_id" bson:"cart_id"`
	}
)
