package model

import "time"

type APPLICATION_TYPE = string

const (
	APPLICATION_TYPE_K8S    APPLICATION_TYPE = "k8s"
	APPLICATION_TYPE_BINARY APPLICATION_TYPE = "binary"
)

// ProductVersion和Application 1对多
type Application struct {
	ApplicationIdentity string           `gorm:"type:varchar(128);not null"`
	ApplicationType     APPLICATION_TYPE `gorm:"type:varchar(128);not null"`
	Owner               string           `gorm:"type:varchar(128);not null"`
	Creator             string           `gorm:"type:varchar(128);not null"`

	// 同wayne中一致
	WayneName          string `gorm:"type:varchar(128);not null"`
	WayneDesc          string `gorm:"type:varchar(128);not null"`
	WayneNamespace     int64  `gorm:"type:bigint(20);not null"`
	WayneApplicationId int64  `gorm:"type:bigint(20);not null"`

	CreateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	UpdateAt *time.Time `gorm:"type:timestamp;not null default current_timestamp"`
	DeleteAt int64      `gorm:"type:bigint(20);not null"`
}
