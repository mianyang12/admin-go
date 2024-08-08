package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"naive-admin-go/common"
	"naive-admin-go/service"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(c *gin.Context) {
	w := c.Writer
	r := c.Request
	//获取路径参数
	categoryTemplate := comnon.Template.Category
	//http://localhost:8080/c/1  1参数 分类的id
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
