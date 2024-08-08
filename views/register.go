package views

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/common"
	"naive-admin-go/config"
)

func (*HTMLApi) Register(c *gin.Context) {
	w := c.Writer
	register := comnon.Template.Login

	register.WriteData(w, config.Cfg.Viewer)
}
