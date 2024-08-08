package dao

import (
	"log"
	"naive-admin-go/model"
)

func UpdatePost(post *model.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}
func DeletePost(user_id, pid int) error {
	_, err := DB.Exec("delete from blog_post where user_id = ? and pid = ?", user_id, pid)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func SavePost(post *model.Post) error {
	// 使用 Exec 方法执行 SQL 插入语句
	ret, err := DB.Exec("INSERT INTO blog_post"+
		"(title, content, markdown, category_id, user_id, view_count, type, slug, create_at, update_at) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)

	if err != nil {
		return err
	}

	// 获取最后插入的 ID
	pid, err := ret.LastInsertId()
	if err != nil {
		return err
	}

	// 设置 Post 的 Pid 字段
	post.Pid = int(pid)
	return nil
}
func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id=?", cId)
	_ = rows.Scan(&count)
	return
}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}
func CountGetAllPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug=?", slug)
	_ = rows.Scan(&count)
	return
}

func GetPostById(pid int) (*model.Post, error) {
	p := &model.Post{}
	err := DB.QueryOne(p, "select * from blog_post where pid=?", pid)
	//p.ViewCount += 1
	//err2, _ := DB.Exec("update blog_post set view_count=? where pid=?", p.ViewCount, pid)
	//if err2 == nil {
	//	return p, nil
	//}
	return p, err
}

// GetPostPage 获取指定页码和页大小的 posts
func GetPostPage(page, pageSize int) ([]model.Post, error) {
	offset := (page - 1) * pageSize
	query := "SELECT * FROM blog_post LIMIT ?, ?"
	return QueryPosts(query, offset, pageSize)
}

// GetPostSearch 搜索包含特定条件的 posts
func GetPostSearch(condition string) ([]model.Post, error) {
	query := "SELECT * FROM blog_post WHERE title LIKE ?"
	return QueryPosts(query, "%"+condition+"%")
}

// GetPostAll 获取所有的 posts
func GetPostAll() ([]model.Post, error) {
	query := "SELECT * FROM blog_post"
	return QueryPosts(query)
}

// GetPostPageByCategoryId 获取指定分类 ID 的 posts 的指定页码和页大小
func GetPostPageByCategoryId(cId, page, pageSize int) ([]model.Post, error) {
	offset := (page - 1) * pageSize
	query := "SELECT * FROM blog_post WHERE category_id = ? LIMIT ?, ?"
	return QueryPosts(query, cId, offset, pageSize)
}

// GetPostPageBySlug 获取指定 Slug 的 posts 的指定页码和页大小
func GetPostPageBySlug(slug string, page, pageSize int) ([]model.Post, error) {
	offset := (page - 1) * pageSize
	query := "SELECT * FROM blog_post WHERE slug = ? LIMIT ?, ?"
	return QueryPosts(query, slug, offset, pageSize)
}
