package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	CategoryID uint      `gorm:"" json:"category_id"`
	Name       string    `gorm:"size:255;not null;" json:"name"`
	Desc       string    `gorm:"type:text;not null;" json:"description"`
	Product    []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:SubCategoryID;associationForeignKey:ID" json:"product"`
}

//----------------------------------SUB CATEGORY--------------------------------------//

func (sc *SubCategory) Prepare() {
	sc.ID = 0
	sc.Name = html.EscapeString(strings.TrimSpace(sc.Name))
	sc.Desc = html.EscapeString(strings.TrimSpace(sc.Desc))
	sc.CreatedAt = time.Now()
	sc.UpdatedAt = time.Now()
}

func (sc *SubCategory) Validate() error {

	if sc.Name == "" {
		return errors.New("required sub category name")
	}
	return nil
}

func (sc *SubCategory) SaveSubCategory(db *gorm.DB) (*SubCategory, error) {
	var err = db.Debug().Model(&SubCategory{}).Create(&sc).Error
	if err != nil {
		return &SubCategory{}, err
	}
	return sc, nil
}

func (sc *SubCategory) FindAllSubCates(db *gorm.DB) (*[]SubCategory, error) {
	var err error
	subCates := []SubCategory{}
	err = db.Debug().Model(&SubCategory{}).Limit(100).Find(&subCates).Error
	if err != nil {
		return &[]SubCategory{}, err
	}
	return &subCates, nil
}

func (sc *SubCategory) FindSubCatesByCateID(db *gorm.DB, pid uint) (*[]SubCategory, error) {
	subCates := []SubCategory{}
	var err = db.Debug().Model(&SubCategory{}).Where("category_id = ?", pid).Find(&subCates).Error
	if err != nil {
		return &subCates, err
	}
	return &subCates, nil
}

func (sc *SubCategory) FindSubCateByID(db *gorm.DB, pid uint) (*SubCategory, error) {
	var err = db.Debug().Model(&SubCategory{}).Where("id = ?", pid).Take(&sc).Error
	if err != nil {
		return &SubCategory{}, err
	}
	//them []Product vao SubCategory
	var product []Product
	db.Debug().Model(&Product{}).Where("sub_category_id = ?", pid).Find(&product)
	for i := 0; i < len(product); i++ {
		sc.Product = append(sc.Product, product[i])
	}

	return sc, nil
}

func (sc *SubCategory) UpdateASubCate(db *gorm.DB) (*SubCategory, error) {

	var err = db.Debug().Model(&SubCategory{}).Where("id = ?", sc.ID).
		Updates(SubCategory{}).Error
	if err != nil {
		return &SubCategory{}, err
	}
	return sc, nil
}

func (sc *SubCategory) DeleteASubCate(db *gorm.DB, pid uint) (int64, error) {
	db = db.Debug().Model(&SubCategory{}).Where("id = ?", pid).Take(&SubCategory{}).Delete(&SubCategory{})
	return db.RowsAffected, nil
}
