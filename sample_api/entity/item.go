package entity

import (
	"errors"
	"sampleApi/db"
	"time"
)

type (
	Item struct {
		ID          int       `gorm:"primarykey;autoIncrement"`
		Name        string    `gorm:"not null"`
		Description string    `gorm:"not null"`
		CreatedAt   time.Time `gorm:"not null"`
	}
	Items []Item
)

// 商品を登録する
func (u *Items) CreateItems() error {
	dbIsntance := db.GetDB()
	return dbIsntance.Create(&u).Error;
}

// 商品IDで商品を取得する
func (u *Items) SelectItemsByItemID(id []int) error {
	dbInstance := db.GetDB()
	if err := dbInstance.Where("id IN ?", id).Find(&u).Error; err != nil {
		return err
	} else if u == nil {
		return errors.New("items not found.")
	}
	return nil
}

func (u *Items) DeleteItemsByItemID(id []int) error {
	dbInstance := db.GetDB()
	return dbInstance.Where("id IN ?", id).Delete(&Item{}).Error;
}
