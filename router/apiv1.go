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

	// 前台登录&操作路由
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
	//================================================================================
	// 后台路由
	auth := r.Group("/auth")
	{
		auth.POST("/login", api.Auth.Login)
		auth.GET("/captcha", api.Auth.Captcha)

		// 添加 JWT 验证中间件
		auth.Use(middleware.Jwt())
		auth.POST("/logout", api.Auth.Logout)
		auth.POST("/password", api.Auth.Password)
	}

	// 用户管理路由
	user := r.Group("/user")
	{
		user.Use(middleware.Jwt())
		user.GET("/", api.User.List)
		user.POST("/", api.User.Add)
		user.DELETE("/:id", api.User.Delete)
		user.PATCH("/:id", api.User.Update)
		user.PATCH("/profile/:id", api.User.Profile)
		user.GET("/detail", api.User.Detail)
	}

	// 角色管理路由
	role := r.Group("/role")
	{
		role.Use(middleware.Jwt())
		role.GET("/", api.Role.List)
		role.POST("/", api.Role.Add)
		role.PATCH("/:id", api.Role.Update)
		role.DELETE("/:id", api.Role.Delete)
		role.PATCH("/users/add/:id", api.Role.AddUser)
		role.PATCH("/users/remove/:id", api.Role.RemoveUser)
		role.GET("/page", api.Role.ListPage)
		role.GET("/permissions/tree", api.Role.PermissionsTree)
	}

	// 权限管理路由
	permission := r.Group("/permission")
	{
		permission.Use(middleware.Jwt())
		permission.POST("/", api.Permissions.Add)
		permission.PATCH("/:id", api.Permissions.PatchPermission)
		permission.DELETE("/:id", api.Permissions.Delete)
		permission.GET("/tree", api.Permissions.List)
		permission.GET("/menu/tree", api.Permissions.List)
	}
}
