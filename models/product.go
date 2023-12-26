package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(300);not null" json:"name"`
	Description string `gorm:"type:text;not null" json:"description"`
	Price       int32  `gorm:"not null" json:"price"`
}
