package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string        `gorm:"size:255;not null;" json:"name"`
	Desc        string        `gorm:"type:text;not null;" json:"description"`
	SubCategory []SubCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID;associationForeignKey:ID" json:"sub_category"`
} //vi du: may tinh, tai nghe, loa,...



func (c *Category) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Desc = html.EscapeString(strings.TrimSpace(c.Desc))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Category) Validate() error {

	if c.Name == "" {
		return errors.New("required category name")
	}
	return nil
}

func (c *Category) SaveCategory(db *gorm.DB) (*Category, error) {
	var err = db.Debug().Model(&Category{}).Create(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) FindAllCategories(db *gorm.DB) (*[]Category, error) {
	var err error
	cates := []Category{}
	err = db.Debug().Model(&Category{}).Limit(100).Find(&cates).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &cates, nil
}

func (c *Category) FindCategoryByID(db *gorm.DB, pid uint) (*Category, error) {
	var err = db.Debug().Model(&Category{}).Where("id = ?", pid).Take(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) UpdateACategory(db *gorm.DB) (*Category, error) {

	var err = db.Debug().Model(&Category{}).Where("id = ?", c.ID).
		Updates(Category{}).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) DeleteACategory(db *gorm.DB, pid uint) (int64, error) {
	db = db.Debug().Model(&Category{}).Where("id = ?", pid).Take(&Category{}).Delete(&Category{})
	return db.RowsAffected, nil
}


