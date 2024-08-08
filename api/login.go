package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"naive-admin-go/common"
	"naive-admin-go/service"
)

func (*Api) Login(c *gin.Context) {
	w := c.Writer
	r := c.Request
	//接收用户名和密码返回对应的json数据
	params := comnon.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	fmt.Println(err)
	if err != nil {
		comnon.Error(w, err)
		return
	}
	comnon.Success(w, loginRes)

}
