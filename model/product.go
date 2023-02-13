package model

import "time"

type Product struct {
	ID              int64  `gorm:"type:bigint(20) unsigned auto_increment;primary_key"`
	ProductIdentity string `gorm:"type:varchar(128);not null"`
	Name            string `gorm:"type:varchar(128);not null"`
	Owner           string `gorm:"type:varchar(128);not null"`
	Desc            string `gorm:"type:varchar(128);not null"`
	Creator         string `gorm:"type:varchar(128);not null"`

	CreateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	UpdateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	DeleteAt int64      `gorm:"type:bigint(20);not null"`
}

// Product和ProductVersion 1对多
type ProductVersion struct {
	ProductVersionIdentity string `gorm:"type:varchar(128);not null"`
	MajorVersion           string `gorm:"type:varchar(128);not null"`
	MinorVersion           string `gorm:"type:varchar(128);not null"`
	Desc                   string `gorm:"type:varchar(128);not null"`
	Creator                string `gorm:"type:varchar(128);not null"`
	BasedStableVersion     string `gorm:"type:varchar(128);not null"`
	IsStableVersion        bool   `gorm:"type:tinyint(1);default 0"`
	IsMajorVersionUpdate   bool   `gorm:"type:tinyint(1);default 0"`

	//todo

	CreateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	UpdateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	DeleteAt int64      `gorm:"type:bigint(20);not null"`

	ProductIdentity string `gorm:"type:varchar(128);not null"`
}
