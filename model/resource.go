package model

import "time"

type Resource struct {
	ResourceIdentity string `gorm:"type:varchar(128);not null"`
	Owner            string `gorm:"type:varchar(128);not null"`
	Creator          string `gorm:"type:varchar(128);not null"`
	Desc             string `gorm:"type:varchar(128);not null"`

	// 同wayne中一致
	WayneName       string `gorm:"type:varchar(128);not null"`
	WayneNamespace  int64  `gorm:"type:bigint(20);not null"`
	WayneResourceId int64  `gorm:"type:bigint(20);not null"`

	ApplicationIdentity string `gorm:"type:varchar(128);not null"`
}

// Application和ResourceVersion 1对多
type ResourceVersion struct {
	ResourceVersionIdentity string `gorm:"type:varchar(128);not null"`
	Version                 string `gorm:"type:varchar(128);not null"`
	Creator                 string `gorm:"type:varchar(128);not null"`
	Desc                    string `gorm:"type:varchar(128);not null"`

	// 同wayne中一致
	WayneResourceVersion int64 `gorm:"type:bigint(20);not null"`

	CreateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	UpdateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	DeleteAt int64      `gorm:"type:bigint(20);not null"`

	ResourceIdentity       string `gorm:"type:varchar(128);not null"`
	ProductVersionIdentity string `gorm:"type:varchar(128);not null"`
}
