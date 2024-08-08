package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Mux(c *gin.Context) {
	w := c.Writer
	r := c.Request
	http.ServeFile(w, r, "template/tupian/index.html")

}
