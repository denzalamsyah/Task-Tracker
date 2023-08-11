package web

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"embed"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

type DosenWeb interface {
	DosenPage(c *gin.Context)
	DosenAddProcess(c *gin.Context)
}

type dosenWeb struct {
	dosenClient     client.DosenClient
	sessionService service.SessionService
	embed          embed.FS
}

func NewDosenWeb(dosenClient client.DosenClient, sessionService service.SessionService, embed embed.FS) *dosenWeb {
	return &dosenWeb{dosenClient, sessionService, embed}
}

func (t *dosenWeb) DosenPage(c *gin.Context) {
	var email string
	if temp, ok := c.Get("email"); ok {
		if contextData, ok := temp.(string); ok {
			email = contextData
		}
	}

	session, err := t.sessionService.GetSessionByEmail(email)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	dosen, err := t.dosenClient.DosenList(session.Token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	var dataTemplate = map[string]interface{}{
		"email": email,
		"dosen": dosen,
	}

	var funcMap = template.FuncMap{
		"exampleFunc": func() int {
			return 0
		},
	}

	var header = path.Join("views", "general", "header.html")
	var filepath = path.Join("views", "main", "dosen.html")

	temp, err := template.New("dosen.html").Funcs(funcMap).ParseFS(t.embed, filepath, header)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	err = temp.Execute(c.Writer, dataTemplate)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
	}
}

func (t *dosenWeb) DosenAddProcess(c *gin.Context) {
	var email string
	if temp, ok := c.Get("email"); ok {
		if contextData, ok := temp.(string); ok {
			email = contextData
		}
	}

	session, err := t.sessionService.GetSessionByEmail(email)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	
	matkulID, _ := strconv.Atoi(c.Request.FormValue("matkul_id"))
	dosen := model.Dosen{
		Name :      c.Request.FormValue("name"),
		Address: c.Request.FormValue("address"),
		MatkulId: matkulID,
	}

	status, err := t.dosenClient.AddDosen(session.Token, dosen)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	if status == 201 {
		c.Redirect(http.StatusSeeOther, "/client/login")
	} else {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message=Add Dosen Failed!")
	}
}
