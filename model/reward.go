// model/reward.go
package model

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	PointCost   int    `json:"point_cost"`
	Quantity    int    `json:"quantity"` // stock available
}
