// model/redemption.go
package model

import "gorm.io/gorm"

type Redemption struct {
	gorm.Model
	MemberID uint   `json:"member_id"`
	RewardID uint   `json:"reward_id"`
	Status   string `json:"status"` // "pending", "approved", "cancelled"
}
