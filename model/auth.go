package model

import (
	datatypes "github.com/nawafilhusnul/NAWNAW-API/common/datatypes"
	"gorm.io/gorm"
)

type Auth struct {
	ID           datatypes.ID         `json:"id" gorm:"column:id;primary_key" example:"$2a$10$/zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3"`
	Email        datatypes.NullString `json:"email" gorm:"column:email" validate:"required,email" example:"email@email.com"`
	Name         datatypes.NullString `json:"name" gorm:"column:name" validate:"required" example:"John Doe"`
	Phone        datatypes.NullString `json:"phone" gorm:"column:phone" validate:"required" example:"+6281234567890"`
	Password     datatypes.HashString `json:"password" gorm:"column:password" validate:"required" example:"* * * *"`
	AccessToken  string               `json:"access_token" gorm:"-" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	RefreshToken string               `json:"refresh_token" gorm:"-" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	Platforms    map[string]bool      `json:"-" gorm:"-" example:"[web:true,mobile:false]"` // use for generate token
	Roles        map[string]bool      `json:"-" gorm:"-" example:"[admin:true,user:false]"` // use for generate token
	Timezone     string               `json:"-" gorm:"-" example:"Asia/Jakarta"`            // use for generate token
}

func (m Auth) TableName() string {
	return "users"
}

type User struct {
	ID          datatypes.ID         `json:"id" gorm:"column:id;primary_key" example:"$2a$10$/zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3zQq3"`
	Email       string               `json:"email" gorm:"column:email" example:"email@email.com"`
	Phone       string               `json:"phone" gorm:"column:phone" example:"+6281234567890"`
	Password    datatypes.HashString `json:"password" gorm:"column:password" example:"* * * *"`
	IsDeleted   bool                 `json:"-" gorm:"column:is_deleted" example:"false"`
	IsActivated bool                 `json:"-" gorm:"column:is_activate" example:"true"`
	DefaultModel
}

func (u User) TableName() string {
	return "users"
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedBy = u.Ctx.GetUser().UserID
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	u.DeletedBy = datatypes.NullInt{Int: u.Ctx.GetUser().UserID, Valid: true}
	return nil
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required" example:"email@email.com"`
	Password   string `json:"password" validate:"required" example:"password"`
	Timezone   string `json:"timezone" validate:"required" example:"Asia/Jakarta"`
}
