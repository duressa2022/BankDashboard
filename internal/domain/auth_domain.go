package domain

// type for working with login request
type LoginRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// type for working with login response
type LoginResponse struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

