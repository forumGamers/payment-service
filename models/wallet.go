package models

import g "github.com/jinzhu/gorm"

type Wallet struct {
	g.Model
	Va_number		string		`json:"va_number" gorm:"NOT NULL"`
	Balance			int			`json:"balance" gorm:"default:0"`
	UserId			int			`json:"userId" gorm:"NOT NULL"`
	Coin			int			`json:"coin" gorm:"default:0"`
}