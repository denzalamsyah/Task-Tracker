package webserv

import (
	"embed"
	"net/http"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
)

type ServerWeb interface {
	IndexServer(c *gin.Context)
}

type serverWeb struct {
	embed embed.FS
}

func NewServerWeb(embed embed.FS) *serverWeb {
	return &serverWeb{embed}
}

func (h *serverWeb) IndexServer(c *gin.Context) {
	var filepath = path.Join("views", "mainserver", "index.html")
	var header = path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(h.embed, filepath, header))

	err := tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}
}
