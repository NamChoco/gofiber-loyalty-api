// model/transaction.go
package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	MemberID uint   `json:"member_id"`
	Type     string `json:"type"` // "earn" or "spend"
	Amount   int    `json:"amount"`
	Note     string `json:"note"`
}
