package database

import (
	"github.com/NamChoco/go-database/model"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	// Seed Members
	members := []model.Member{
		{Name: "Alice", Email: "alice@example.com", Phone: "0800000001", Points: 100},
		{Name: "Bob", Email: "bob@example.com", Phone: "0800000002", Points: 150},
		{Name: "Charlie", Email: "charlie@example.com", Phone: "0800000003", Points: 50},
	}
	for _, m := range members {
		db.Create(&m)
	}

	// Seed Rewards
	rewards := []model.Reward{
		{Name: "Starbucks Voucher", Description: "100 THB Gift Card", PointCost: 100, Quantity: 10},
		{Name: "Movie Ticket", Description: "Major Cineplex Voucher", PointCost: 80, Quantity: 5},
		{Name: "Wireless Earbuds", Description: "Bluetooth earbuds", PointCost: 300, Quantity: 2},
	}
	for _, r := range rewards {
		db.Create(&r)
	}

	// Seed Transactions
	tx := []model.Transaction{
		{MemberID: 1, Type: "earn", Amount: 50, Note: "Signup Bonus"},
		{MemberID: 1, Type: "earn", Amount: 50, Note: "Purchase"},
		{MemberID: 2, Type: "earn", Amount: 100, Note: "Referral"},
		{MemberID: 2, Type: "spend", Amount: 50, Note: "Redeem coffee"},
		{MemberID: 3, Type: "earn", Amount: 50, Note: "Promo"},
	}
	for _, t := range tx {
		db.Create(&t)
	}

	// Seed Redemptions
	redemptions := []model.Redemption{
		{MemberID: 1, RewardID: 1, Status: "approved"},
		{MemberID: 2, RewardID: 2, Status: "pending"},
	}
	for _, r := range redemptions {
		db.Create(&r)
	}
}
