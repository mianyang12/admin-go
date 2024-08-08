package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"naive-admin-go/api"
	"naive-admin-go/middleware"
	"naive-admin-go/utils"
	"naive-admin-go/views"
)

func Init(r *gin.Engine) {
	// 使用 cookie 存储会话数据
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	// CORS 设置
	r.Use(middleware.Cors())
	// 静态文件服务
	r.Static("/resource", "./public/resource")
	r.Static("/p/resource", "./public/resource")

	// 前台路由
	view := r.Group("/")
	{
		view.GET("", views.HTML.Index)
		view.GET("c/", views.HTML.Category)
		view.GET("login", views.HTML.Login)
		view.GET("/p/:pid", utils.ViewCount(), views.HTML.Detail)
		view.GET("writing", views.HTML.Writing)
		view.GET("pigeonhole", views.HTML.Pigeonhole)
		view.GET("picture", views.Mux)
		view.GET("p/picture", views.Mux)
	}

	// API 路由
	a := r.Group("/api/v1")
	{
		a.POST("/login", api.API.Login)

		// 除了登录不需要鉴权
		a.Use(utils.AuthJwt())
		a.PUT("/post", api.API.SaveAndUpdatePost)
		a.POST("/post/", api.API.SaveAndUpdatePost)
		a.GET("/post/:pid", api.API.GetPost)
		a.POST("post/search", api.API.SearchPost)
		a.POST("/qiniu/token", api.API.QiniuToken)
		a.DELETE("/post/delete/:pid", api.API.DeletePost)
	}

	// 后台路由
	r.POST("/auth/login", api.Auth.Login)
	r.GET("/auth/captcha", api.Auth.Captcha)

	r.Use(middleware.Jwt())
	r.POST("/auth/logout", api.Auth.Logout)
	r.POST("/auth/password", api.Auth.Logout)

	r.GET("/user", api.User.List)
	r.POST("/user", api.User.Add)
	r.DELETE("/user/:id", api.User.Delete)
	r.PATCH("/user/password/reset/:id", api.User.Update)
	r.PATCH("/user/:id", api.User.Update)
	r.PATCH("/user/profile/:id", api.User.Profile)
	r.GET("/user/detail", api.User.Detail)

	r.GET("/role", api.Role.List)
	r.POST("/role", api.Role.Add)
	r.PATCH("/role/:id", api.Role.Update)
	r.DELETE("/role/:id", api.Role.Delete)
	r.PATCH("/role/users/add/:id", api.Role.AddUser)
	r.PATCH("/role/users/remove/:id", api.Role.RemoveUser)
	r.GET("/role/page", api.Role.ListPage)
	r.GET("/role/permissions/tree", api.Role.PermissionsTree)

	r.POST("/permission", api.Permissions.Add)
	r.PATCH("/permission/:id", api.Permissions.PatchPermission)
	r.DELETE("/permission/:id", api.Permissions.Delete)
	r.GET("/permission/tree", api.Permissions.List)
	r.GET("/permission/menu/tree", api.Permissions.List)
}
