package dao

import (
	"fmt"
	"log"
	"naive-admin-go/model"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select username from blog_user where id=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, password string) *model.User {
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
	log.Println("", row, username, password)
	var user = &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println(user)
	return user // 返回用户信息
}
