package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"naive-admin-go/common"
	"naive-admin-go/dao"
	"naive-admin-go/model"
	"naive-admin-go/service"
	"naive-admin-go/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(c *gin.Context) {
	w := c.Writer
	r := c.Request
	//获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)

	if err != nil {
		comnon.Error(w, errors.New("不识别此请求路径"))
		return
	}

	post, err := dao.GetPostById(pid)
	if err != nil {
		comnon.Error(w, err)
		return
	}

	comnon.Success(w, post)
}
func (*Api) DeletePost(c *gin.Context) {
	//获取用户id，判断用户是否登录
	w := c.Writer
	r := c.Request
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		comnon.Error(w, errors.New("登录已过期"))
		return
	}
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/delete/")
	pid, err := strconv.Atoi(pIdStr)
	userId := claim.Uid
	err = dao.DeletePost(userId, pid)
	if err != nil {
		comnon.Error(w, err)
		return
	}
	comnon.Success(w, nil)

}

func (*Api) SaveAndUpdatePost(c *gin.Context) {
	w := c.Writer
	r := c.Request

	// 获取用户id，判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		comnon.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid

	// POST save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := comnon.GetRequestJsonParam(r)

		var categoryId int
		if cId, ok := params["categoryId"].(string); ok {
			var err error
			categoryId, err = strconv.Atoi(cId)
			if err != nil {
				comnon.Error(w, errors.New("无效的 categoryId"))
				return
			}
		} else {
			comnon.Error(w, errors.New("categoryId 不存在"))
			return
		}

		content, ok := params["content"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 content"))
			return
		}

		markdown, ok := params["markdown"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 markdown"))
			return
		}

		slug, ok := params["slug"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 slug"))
			return
		}

		title, ok := params["title"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 title"))
			return
		}

		postType, ok := params["type"].(float64)
		if !ok {
			comnon.Error(w, errors.New("无效的 type"))
			return
		}
		pType := int(postType)

		post := &model.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		comnon.Success(w, post)

	case http.MethodPut:
		// Update
		params := comnon.GetRequestJsonParam(r)

		var categoryId int
		if cId, ok := params["categoryId"].(string); ok {
			var err error
			categoryId, err = strconv.Atoi(cId)
			if err != nil {
				comnon.Error(w, errors.New("无效的 categoryId"))
				return
			}
		} else {
			comnon.Error(w, errors.New("categoryId 不存在"))
			return
		}

		content, ok := params["content"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 content"))
			return
		}

		markdown, ok := params["markdown"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 markdown"))
			return
		}

		slug, ok := params["slug"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 slug"))
			return
		}

		title, ok := params["title"].(string)
		if !ok {
			comnon.Error(w, errors.New("无效的 title"))
			return
		}

		postType, ok := params["type"].(float64)
		if !ok {
			comnon.Error(w, errors.New("无效的 type"))
			return
		}
		pType := int(postType)

		pidFloat, ok := params["pid"].(float64)
		if !ok {
			comnon.Error(w, errors.New("无效的 pid"))
			return
		}
		pid := int(pidFloat)

		post := &model.Post{
			Pid:        pid,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		comnon.Success(w, post)
	}
}
func (*Api) SearchPost(c *gin.Context) {
	w := c.Writer
	r := c.Request
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	comnon.Success(w, searchResp)
}
