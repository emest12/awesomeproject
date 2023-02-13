package app

import "time"

type Product struct {
	ProductIdentity string
	Name            string
	Owner           string
	Desc            string

	CreateAt *time.Time
	UpdateAt *time.Time
}

// Product和ProductVersion 1对多
type ProductVersion struct {
	ProductVersionIdentity string
	MajorVersion           string
	MinorVersion           string
	Desc                   string
	Creator                string
	IsStableVersion        bool

	//todo

	CreateAt *time.Time
	UpdateAt *time.Time

	ProductIdentity string
}
