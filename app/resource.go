package app

import "time"

// Application和Resource 1对多
type Resource struct {
	ResourceIdentity string
	Version          string

	// 同wayne中一致
	Name                 string
	WayneApplicationId   int64
	WayneNamespace       int64
	WayneResourceVersion int64

	CreateAt *time.Time
	UpdateAt *time.Time

	ApplicationIdentity string
}
