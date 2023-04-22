package sharedsecret

type SharedSecret struct {
	ClientID     string `json:"client_id" binding:"required"`
	SharedSecret string `json:"shared_secret" binding:"required"`
}
