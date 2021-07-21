package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string          `gorm:"size:255;not null;unique" json:"name"`
	ImageCover    string          `gorm:"size:255;not null;" json:"image_cover"`
	PriceCover    string          `gorm:"size:255;not null;" json:"price_cover"`
	Desc          string          `gorm:"size:255;type:text;not null;" json:"description"`
	BrandID       uint            `gorm:"" json:"brand_id"`
	SubCategoryID uint            `gorm:"" json:"sub_category_id"`
	Variant       []Variant       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ProductID;" json:"variant"`
	ProductOption []ProductOption `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ProductID;" json:"product_option"`
}

type ProductOption struct {
	gorm.Model
	ProductID uint `gorm:"" json:"product_id"`
	OptionID  uint `gorm:"" json:"option_id"`
}

func (p *Product) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Desc = html.EscapeString(strings.TrimSpace(p.Desc))
	p.ImageCover = html.EscapeString(strings.TrimSpace(p.ImageCover))
	p.PriceCover = html.EscapeString(strings.TrimSpace(p.PriceCover))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("required product name")
	}
	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint) (*Product, error) {
	var err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) FindProductsByName(db *gorm.DB, pname string) (*[]Product, error) {
	products := []Product{}
	// var err = db.Debug().Exec("SELECT * FROM 'products' WHERE 'products'.name like ? AND 'products'.deleted_at IS NULL LIMIT 100", pname).Scan(products).Error
	var err = db.Debug().Model(&Product{}).Limit(100).Where("name like ?", "%"+pname+"%").Find(&products).Error
	if err != nil {
		return &products, err
	}
	return &products, nil
}

func (p *Product) FindProductsBySubCategoryID(db *gorm.DB, pid uint) (*[]Product, error) {
	products := []Product{}
	var err = db.Debug().Model(&Product{}).Where("sub_category_id = ?", pid).Find(&products).Error
	if err != nil {
		return &products, err
	}
	return &products, nil
}

func (p *Product) FindProductsByBrandID(db *gorm.DB, pid uint) (*[]Product, error) {
	products := []Product{}
	var err = db.Debug().Model(&Product{}).Where("brand_id = ?", pid).Find(&products).Error
	if err != nil {
		return &products, err
	}
	return &products, nil
}

func (p *Product) UpdateAProduct(db *gorm.DB) (*Product, error) {

	var err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).
		Updates(Product{Name: p.Name, ImageCover: p.ImageCover, PriceCover: p.PriceCover, Desc: p.Desc, BrandID: p.BrandID, SubCategoryID: p.SubCategoryID}).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteAProduct(db *gorm.DB, pid uint) (int64, error) {
	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).Delete(&Product{})
	return db.RowsAffected, nil
}
