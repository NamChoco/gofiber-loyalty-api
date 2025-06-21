// model/member.go
package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Phone     string         `json:"phone"`
	Points    int            `json:"points"`
	Transactions []Transaction `gorm:"foreignKey:MemberID"`
	Redemptions  []Redemption  `gorm:"foreignKey:MemberID"`
}
