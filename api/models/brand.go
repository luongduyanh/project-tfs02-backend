package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name string `gorm:"size:255;not null;" json:"name"`
	Desc string `gorm:"type:text;not null;" json:"description"`

	Product []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:BrandID;associationForeignKey:ID" json:"product"`
} //vi du: asus, dell, samsung,...

func (b *Brand) Prepare() {
	b.ID = 0
	b.Name = html.EscapeString(strings.TrimSpace(b.Name))
	b.Desc = html.EscapeString(strings.TrimSpace(b.Desc))
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Brand) Validate() error {
	if b.Name == "" {
		return errors.New("required brand name")
	}
	return nil
}

func (b *Brand) SaveBrand(db *gorm.DB) (*Brand, error) {
	var err = db.Debug().Model(&Product{}).Create(&b).Error
	if err != nil {
		return &Brand{}, err
	}
	return b, nil
}

func (b *Brand) FindAllBrands(db *gorm.DB) (*[]Brand, error) {
	//trả về tất cả các brand nhưng không gồm các product
	// var err error
	// brands := []Brand{}
	// err = db.Debug().Model(&Brand{}).Limit(100).Find(&brands).Error
	// if err != nil {
	// 	return &[]Brand{}, err
	// }

	//trả về cả các product của từng brand
	var listBrands []Brand
	db.Debug().Model(&listBrands).Limit(100).Find(&listBrands)
	for i, brand := range listBrands {
		var listProducts []Product
		db.Where("brand_id = ?", brand.ID).Find(&listProducts)
		listBrands[i].Product = listProducts
	}
	return &listBrands, nil
}

func (b *Brand) FindBrandByID(db *gorm.DB, pid uint) (*Brand, error) {
	var err = db.Debug().Model(&Brand{}).Where("id = ?", pid).Take(&b).Error
	if err != nil {
		return &Brand{}, err
	}
	//them []Product vao Brand
	var product []Product
	db.Debug().Where("brand_id = ?", pid).Find(&product)
	for i := 0; i < len(product); i++ {
		b.Product = append(b.Product, product[i])
	}
	return b, nil
}

func (b *Brand) UpdateABrand(db *gorm.DB) (*Brand, error) {
	var err = db.Debug().Model(&Brand{}).Where("id = ?", b.ID).
		Updates(Brand{}).Error
	if err != nil {
		return &Brand{}, err
	}
	return b, nil
}

func (b *Brand) DeleteABrand(db *gorm.DB, pid uint) (int64, error) {
	db = db.Debug().Model(&Brand{}).Where("id = ?", pid).Take(&Brand{}).Delete(&Brand{})
	return db.RowsAffected, nil
}
