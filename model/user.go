package model

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Enable     bool      `json:"enable"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime"`
}
type Users struct {
	Id         int       `json:"id"`
	UserName   string    `json:"username"`
	Password   string    `json:"password"`
	Avatar     string    `json:"avatar"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}

func (User) TableName() string {
	return "user"
}
