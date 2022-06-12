package entity

import "time"

type Campaign struct {
	Id            int64
	Title         string
	Description   string
	TargetAmount  int
	CurrentAmount int
	TimeStart     time.Time
	TimeEnd       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CampaignImage struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
