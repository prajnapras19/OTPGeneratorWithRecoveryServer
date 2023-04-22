package sharedsecret

import (
	"time"
)

type SharedSecret struct {
	ClientID     string `json:"client_id" binding:"required"`
	SharedSecret string `json:"shared_secret" binding:"required"`
}

type SharedSecretInDatabase struct {
	SharedSecret
	CreatedAt time.Time `json:"created_at"`
}
