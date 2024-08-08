package dao

import (
	"fmt"
	"log"
	"naive-admin-go/model"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("SELECT username FROM blog_user WHERE id=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, password string) *model.Users {
	// 查询数据库
	row := DB.QueryRow(
		"SELECT * FROM blog_user WHERE username = ? AND password = ? ",
		username,
		password,
	)

	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var users = &model.Users{}
	err := row.Scan(&users.Id, &users.UserName, &users.Password, &users.Avatar, &users.CreateTime, &users.UpdateTime)
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println(users)
	return users // 返回用户信息
}
