package model

import "time"

type Order struct {
	OrderId      uint      `gorm:"primary_key" json:"order_id"`
	CustomerName string    `gorm:"not null" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemId      uint   `gorm:"primary_key" json:"item_id"`
	ItemCode    string `gorm:"not null" json:"item_code"`
	Description string `gorm:"not null;type:varchar(20)" json:"desc"`
	Quantity    int8   `gorm:"not null" json:"qty"`
	OrderID     uint
}
