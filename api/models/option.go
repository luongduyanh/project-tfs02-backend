package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Option        string          `gorm:"size:255;not null;" json:"option"`
	ProductOption []ProductOption `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:OptionID;" json:"product_option"`
	OptionValue   []OptionValue   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:OptionID;" json:"option_value"`
} // mau sac, kich thuoc, ram, cpu,...

type OptionValue struct {
	gorm.Model
	OptionID  uint   `gorm:"" json:"option_id"`
	VariantID uint   `gorm:"" json:"variant_id"`
	Value     string `gorm:"size:255;not null;" json:"value"`
} // do, vang, 2gb, 4gb, i5, i7,...
	