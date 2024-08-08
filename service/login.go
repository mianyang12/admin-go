package service

import (
	"errors"
	"naive-admin-go/dao"
	"naive-admin-go/model"
	"naive-admin-go/utils"
)

func Login(username, password string) (*model.LoginRes, error) {
	password = utils.Md5Crypt(password, "yang")
	users := dao.GetUser(username, password)
	if users == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := users.Id
	//生成token  jwt技术进行生成 令牌  A.B.C
	token := utils.GenerateToken(uid)
	var userInfo model.UserInfo
	userInfo.Uid = users.Id
	userInfo.UserName = users.UserName
	var lr = &model.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
