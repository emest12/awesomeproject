package app

import "time"

type APPLICATION_TYPE = string

const (
	APPLICATION_TYPE_K8S    APPLICATION_TYPE = "k8s"
	APPLICATION_TYPE_BINARY APPLICATION_TYPE = "binary"
)

// ProductVersion和Application 1对多
type Application struct {
	ApplicationIdentity string
	ApplicationType     APPLICATION_TYPE

	// 同wayne中一致
	Name               string
	Desc               string
	WayneApplicationId int64
	WayneNamespace     int64

	CreateAt *time.Time
	UpdateAt *time.Time

	ProductVersionIdentity string
}
