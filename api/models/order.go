package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"" json:"user_id"`
	StatusID   uint        `gorm:"not null;default:0" json:"status_id"`
	TotalPrice string      `gorm:"" json:"total_price"`
	OrderLine  []OrderLine `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:OrderID;" json:"order_line"`
	Ship       Ship        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:OrderID;" json:"ship"`
}

type OrderLine struct {
	gorm.Model
	OrderID      uint   `gorm:"" json:"order_id"`
	ProductName  string `gorm:"size:255;not null" json:"product_name"`
	ProductPrice string `gorm:"size:255;not null" json:"product_price"`
	Quantity     uint   `gorm:"not null" json:"quatity"`
	TotalPrice   string `gorm:"" json:"total_price"`
}

func (o *Order) SaveOrder(db *gorm.DB) (*Order, error) {
	var err = db.Debug().Model(&Order{}).Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}

func (o *Order) FindAllOrders(db *gorm.DB) (*[]Order, error) {

	orders := []Order{}
	var err = db.Debug().Model(&Order{}).Limit(100).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	return &orders, nil
}

func (o *Order) FindOrderByID(db *gorm.DB, pid uint64) (*Order, error) {
	var err = db.Debug().Model(&Order{}).Where("id = ?", pid).Take(&o).Error
	if err != nil {
		return &Order{}, err
	}

	return o, nil
}

func (o *Order) FindOrderLinesByOrderID(db *gorm.DB, pid uint) (*[]OrderLine, error) {
	orderLines := []OrderLine{}
	var err = db.Debug().Model(&OrderLine{}).Where("order_id = ?", pid).Find(&orderLines).Error
	if err != nil {
		return &orderLines, err
	}
	return &orderLines, nil
}

func (o *Order) UpdateAOrder(db *gorm.DB) (*Order, error) {

	var err = db.Debug().Model(&Order{}).Where("id = ?", o.ID).Updates(Order{StatusID: o.StatusID}).Error
	if err != nil {
		return &Order{}, err
	}

	return o, nil
}

func (o *Order) DeleteAOrder(db *gorm.DB, pid uint, uid uint) (int64, error) {

	db = db.Debug().Model(&Order{}).Where("id = ?", pid, uid).Take(&Order{}).Delete(&Order{})
	return db.RowsAffected, nil
}
