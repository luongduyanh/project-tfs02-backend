package models

import "gorm.io/gorm"

type Ship struct {
	gorm.Model
	OrderID         uint   `gorm:"" json:"order_id"`
	NameOfCutomer   string `gorm:"size:255;not null" json:"name_of_cutomer"`
	PhoneOfCustomer string `gorm:"size:255;not null" json:"phone_of_customer"`
	EmailOfCustomer string `gorm:"size:255;not null" json:"email_of_customer"`
	City            string `gorm:"size:255;not null" json:"city"`
	District        string `gorm:"size:255;not null" json:"district"`
	Home            string `gorm:"size:255;not null" json:"home"`
	Note            string `gorm:"size:255;" json:"note"`
	Price           string `gorm:"not null;" json:"price"`
}
