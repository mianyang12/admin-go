package service

import (
	"html/template"
	"log"
	"naive-admin-go/config"
	"naive-admin-go/dao"
	"naive-admin-go/model"
)

func SavePost(post *model.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *model.Post) {
	dao.UpdatePost(post)
}
func GetPostDetail(pid int) (*model.PostRes, error) {
	post, err := dao.GetPostById(pid)
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)
	postMore := model.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     model.DateDay(post.CreateAt),
		UpdateAt:     model.DateDay(post.UpdateAt),
	}
	var postRes = &model.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	return postRes, nil
}

func Writing() (wr model.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}

func SearchPost(condition string) []model.SearchResp {
	posts, _ := dao.GetPostSearch(condition)
	var searchResps []model.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, model.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return searchResps
}
