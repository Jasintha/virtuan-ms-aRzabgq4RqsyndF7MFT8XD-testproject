package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	Model       `bson:"-"`
	MID_        string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Address     string `bson:"address" json:"address" xml:"address"`
	Age         int64  `bson:"age" json:"age" xml:"age"`
	Contact     string `bson:"contact" json:"contact" xml:"contact"`
	Email       string `bson:"email" json:"email" xml:"email"`
	Name        string `bson:"name" json:"name" xml:"name"`
	Nationality string `bson:"nationality" json:"nationality" xml:"nationality"`
	Nic         string `bson:"nic" json:"nic" xml:"nic"`
}

func (Customer) TableName() string {
	return "customer"
}
func (m *Customer) PreloadCustomer(db *gorm.DB) *gorm.DB {
	return db
}
