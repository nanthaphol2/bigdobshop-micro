package auth

import (
	"time"

	"github.com/nanthaphol2/bigdobshop-micro/modules/profile"
)

type (
	ProfileLoginReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	RefreshTokenReq struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required,max=500"`
	}

	InsertProfileRole struct {
		ProfileId string `json:"profile_id" validate:"required"`
		RoleCode  []int  `json:"role_id" validate:"required"`
	}

	ProfileIntercepter struct {
		*profile.ProfileProfile
		Credential *CredentialRes `json:"credential"`
	}

	CredentialRes struct {
		Id           string    `json:"_id" bson:"_id,omitempty"`
		ProfileId    string    `json:"profile_id" bson:"profile_id"`
		RoleCode     int       `json:"role_code" bson:"role_code"`
		AccessToken  string    `json:"access_token" bson:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)