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
	Uid      int       `json:"uid"`
	UserName string    `json:"userName"`
	Passwd   string    `json:"passwd"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}

func (User) TableName() string {
	return "user"
}
