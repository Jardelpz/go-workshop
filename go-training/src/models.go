package src

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type user struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Debts     []debt    `gorm:"foreignkey:UserRefer;association_foreignkey:ID" json:"debts"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"type:varchar(100); not null; unique_index" json:"email"`
	BirthDate string    `gorm:"type:varchar(20)" json:"birth_date"`
	CreatedAt time.Time `gorm:"default: NOW()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default: NOW()" json:"updated_at"`
}

func (user) TableName() string {
	return "users"
}

type debt struct {
	ID           string    `gorm:"primary_key" json:"id"`
	UserRefer    int       //`gorm:json:"user_id"`
	Company_name string    `gorm:"type:varchar(45); not null" json:"company_name"`
	Value        float64   `gorm:"not null" json:"value"`
	Date         string    `gorm:"not null" json:"date"`
	Status       int       `gorm:"not null" json:"status"`
	CreatedAt    time.Time `gorm:"default: NOW()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default: NOW()" json:"updated_at"`
}

func (debt) TableName() string {
	return "debts"
}

//BeforeCreate execute before insert data on DB
func (o *debt) BeforeCreate(scope *gorm.Scope) error {

	uuid := uuid.NewV4().String()
	return scope.SetColumn("ID", uuid)
}
