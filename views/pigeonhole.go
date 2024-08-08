package views

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/common"
	"naive-admin-go/service"
)

func (*HTMLApi) Pigeonhole(c *gin.Context) {
	w := c.Writer
	pigeonhole := comnon.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
