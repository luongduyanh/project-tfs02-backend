package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ImgSrc    string `gorm:"size:255;not null;" json:"src"`
	VariantID uint   `gorm:"" json:"variant_id"`
}

func (i *Image) SaveImage(db *gorm.DB) (*Image, error) {
	var err = db.Debug().Model(&Image{}).Create(&i).Error
	if err != nil {
		return &Image{}, err
	}
	return i, nil
}

func (i *Image) FindAllImages(db *gorm.DB) (*[]Image, error) {
	var err error
	images := []Image{}
	err = db.Debug().Model(&Post{}).Limit(100).Find(&images).Error
	if err != nil {
		return &[]Image{}, err
	}
	return &[]Image{}, nil
}

func (i *Image) FindImageByID(db *gorm.DB, pid uint) (*Image, error) {
	var err = db.Debug().Model(&Image{}).Where("id = ?", pid).Take(&i).Error
	if err != nil {
		return &Image{}, err
	}
	return i, nil
}

func (i *Image) UpdateAImage(db *gorm.DB) (*Image, error) {

	var err = db.Debug().Model(&Image{}).Where("id = ?", i.ID).
		Updates(Image{ImgSrc: i.ImgSrc, VariantID: i.VariantID}).Error
	if err != nil {
		return &Image{}, err
	}
	return i, nil
}

func (i *Image) DeleteAImage(db *gorm.DB, pid uint) (int64, error) {
	db = db.Debug().Model(&Image{}).Where("id = ?", pid).Take(&Image{}).Delete(&Image{})
	return db.RowsAffected, nil
}
