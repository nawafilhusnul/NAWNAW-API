package model

import (
	datatypes "github.com/nawafilhusnul/big-app/common/datatypes"
)

type Auth struct {
	ID           datatypes.ID `json:"id" gorm:"column:id;primary_key" example:"$2a$10$/zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3"`
	Email        string       `json:"email" gorm:"column:email" example:"email@email.com"`
	Phone        string       `json:"phone" gorm:"column:phone" example:"+6281234567890"`
	Password     string       `json:"-" gorm:"column:password" example:"* * * *"`
	AccessToken  string       `json:"access_token" gorm:"-" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	RefreshToken string       `json:"refresh_token" gorm:"-" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

func (m Auth) TableName() string {
	return "users"
}
