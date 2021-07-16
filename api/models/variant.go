package models

import "gorm.io/gorm"

type Variant struct {
	gorm.Model
	SKU         string        `gorm:"size:255;not null;" json:"sku"` // dinh danh duy nhat 1 product
	VariantName string        `gorm:"size:255;not null;" json:"variant_name"`
	Desc        string        `gorm:"not null;type:text" json:"description"`
	Price       string        `gorm:"not null;" json:"price"`
	ProductID   uint          `json:"product_id"`
	Image       []Image       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:VariantID;" json:"image"`
	OptionValue []OptionValue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:VariantID;" json:"option_value"`
}


