package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"naive-admin-go/common"
	"naive-admin-go/service"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(c *gin.Context) {
	w := c.Writer
	r := c.Request
	detail := comnon.Template.Detail
	//获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)

}
