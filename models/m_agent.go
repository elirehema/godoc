package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agent struct {
	Id         string `gorm:"column:id;not null unique;<-:create; primarykey" json:"id"`
	Phone      string `db:"phone" gorm:"unique" json:"phone"`
	Email      string `db:"email" gorm:"unique" json:"email"`
	FirstName  string `db:"firstnname" json:"firstname"`
	LastName   string `db:"lastname" json:"lastname"`
	AgentCode  string `gorm:"not null;not null;unique;<-:create" json:"agentCode"`
	DistrictId string `gorm:"column:district_id;size:50; not null" json:"districtId"`
	IsActive   bool   `db:"is_active" gorm:"default:true;not null" json:"isActive"`
	UserId     string `gorm:"column:user_id; size:50; not null" json:"userId"`
}

func (d *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	d.Id = uuid.New().String()
	return
}
