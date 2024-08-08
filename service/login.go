package service

import (
	"errors"
	"naive-admin-go/dao"
	"naive-admin-go/model"
	"naive-admin-go/utils"
)

func Login(username, password string) (*model.LoginRes, error) {
	password = utils.Md5Crypt(password, "yang")
	user := dao.GetUser(username, password)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.ID
	//生成token  jwt技术进行生成 令牌  A.B.C
	token := utils.GenerateToken(uid)
	var userInfo model.UserInfo
	userInfo.Uid = user.ID
	userInfo.UserName = user.Username
	userInfo.Avatar = ""
	var lr = &model.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
