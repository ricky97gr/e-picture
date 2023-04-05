package uimodel

import "my-admin/pkg/time"

type User struct {
	Name       string          `json:"name"`
	Role       string          `json:"role"`
	ID         string          `json:"id"`
	Phone      string          `json:"phone"`
	Email      string          `json:"email"`
	CreateTime *time.LocalTime `json:"createTime"`
}
