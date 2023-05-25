package models

type CredentialRequestDTO struct {
	PlatformName string `json:"platformName"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
}

type UserEntity struct {
	Id     string
	Mobile int64
}

type CredentialEntity struct {
	Id           string `json:"id"`
	PlatformName string `json:"platformName"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	CreatedAt    string `json:"createdAt"`
	ModifiedAt   string `json:"modifiedAt"`
}
