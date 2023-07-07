package model

type Product struct {
	Id    int     `gorm:"type:int;primary_key"`
	Name  string  `gorm:"not null" json:"name"`
	Price float64 `gorm:"not null" json:"price"`
	Stock int     `gorm:"not null" json:"stock"`
}
