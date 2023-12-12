package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Credential struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		ProfileId   string             `json:"profile_id" bson:"profile_id"`
		RoleCode    int                `json:"role_code" bson:"role_code"`
		AccessToken string             `json:"access_token" bson:"access_token"`
		CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt   time.Time          `json:"update_at" bson:"updated_at"`
	}

	Role struct {
		Id    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title string             `json:"title" bson:"title"`
		Code  int                `json:"code" bson:"code"`
	}

	// UpdateRefreshTokenReq struct {
	// 	ProfileId    string    `bson:"profile_id"`
	// 	AccessToken  string    `bson:"access_token"`
	// 	RefreshToken string    `bson:"refresh_token"`
	// 	UpdatedAt    time.Time `bson:"updated_at"`
	// }
)
