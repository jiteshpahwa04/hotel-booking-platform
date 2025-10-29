package models

import "time"

type Role struct {
	Id          int64 		`json:"id"`
	Name        string		`json:"name"`
	Description string		`json:"description"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt   time.Time	`json:"updated_at"`
}

type Permission struct {
	Id          int64
	Name        string
	Description string
	Resource    string
	Action      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RolePermission struct {
	Id           int64
	RoleId       int64
	PermissionId int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRole struct {
	Id        int64
	UserId    int64
	RoleId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
