package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"naive-admin-go/dao"
	"naive-admin-go/model"
	"net/http"
	"strconv"
	"strings"
)

// ViewCount 是一个用于增加文章浏览计数的中间件。
func ViewCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取文章ID (pid)，这里假设它是URL的一部分。
		pidStr := c.Param("pid")
		pidStr = strings.TrimSuffix(pidStr, ".html") // 移除扩展名
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			log.Println("Failed to convert PID to integer:", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		p := &model.Post{}
		err = dao.DB.QueryOne(p, "SELECT * FROM blog_post WHERE pid=?", pid)
		if err != nil {
			log.Println("Failed to fetch post:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		p.ViewCount++
		//更新对应的viewCount
		_, err = dao.DB.Exec("UPDATE blog_post SET view_count=? WHERE pid=?", p.ViewCount, pid)
		if err != nil {
			log.Println("Failed to update view count:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

// AuthJwt 是鉴权中间件
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的认证令牌
		tokenString := c.GetHeader("Authorization")

		// 解析JWT令牌
		token, claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录已过期"})
			c.Abort()
			return
		}

		// 检查令牌是否有效
		if token != nil && claims != nil && token.Valid {
			c.Set("uid", claims.Uid)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录已过期"})
			c.Abort()
			return
		}
	}
}
