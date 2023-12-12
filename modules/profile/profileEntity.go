package profile

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Profile struct {
		Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email        string             `json:"email" bson:"email"`
		Password     string             `json:"password" bson:"password"`
		Username     string             `json:"username" bson:"username"`
		CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
		ProfileRoles []ProfileRole      `bson:"profile_roles"`
	}

	ProfileRole struct {
		RoleTitle string `json:"role_title" bson:"role_title"`
		RoleCode  int    `json:"role_code" bson:"role_code"`
	}

	ProfileProfileBson struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Email     string             `json:"email" bson:"email"`
		Username  string             `json:"username" bson:"username"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	}

	ProfileSavingAccount struct {
		ProfileId string  `json:"profile_id" bson:"profile_id"`
		Balance   float64 `json:"balance" bson:"balance"`
	}

	ProfileTransaction struct {
		Id        primitive.ObjectID `bson:"_id,omitempty"`
		ProfileId string             `bson:"profile_id"`
		Amount    float64            `bson:"amount"`
		CreatedAt time.Time          `bson:"created_at"`
	}
)
