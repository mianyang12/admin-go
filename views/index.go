package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"naive-admin-go/common"
	"naive-admin-go/context"
	"naive-admin-go/service"
	"strconv"
	"strings"
)

func (*HTMLApi) IndexTest(ctx *context.MsContext) {
	log.Println(ctx.GetPathVariable("id"))

}

func (*HTMLApi) Index(c *gin.Context) {
	w := c.Writer
	r := c.Request
	index := comnon.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	//每页显示的数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误!请联系管理员"))
	}
	index.WriteData(w, hr)
}
