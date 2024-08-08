package service

import (
	"html/template"
	"naive-admin-go/config"
	"naive-admin-go/dao"
	"naive-admin-go/model"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*model.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts []model.Post
	var total int
	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize)
		total = dao.CountGetAllPost()
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)
		total = dao.CountGetAllPostBySlug(slug)
	}

	var postMores []model.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := model.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			model.DateDay(post.CreateAt),
			model.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	//11  10 2  10 1 9 1  21 3
	//  (11-1)/10 + 1 = 2

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &model.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	return hr, nil
}
