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

type StudentWeb interface {
	StudentPage(c *gin.Context)
	StudentAddProcess(c *gin.Context)
}

type studentWeb struct {
	studentClient     client.StudentClient
	sessionService service.SessionService
	embed          embed.FS
}

func NewStudentWeb(studentClient client.StudentClient, sessionService service.SessionService, embed embed.FS) *studentWeb {
	return &studentWeb{studentClient, sessionService, embed}
}

func (t *studentWeb) StudentPage(c *gin.Context) {
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

	student, err := t.studentClient.StudentList(session.Token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	var dataTemplate = map[string]interface{}{
		"email": email,
		"student": student,
	}

	var funcMap = template.FuncMap{
		"exampleFunc": func() int {
			return 0
		},
	}

	var header = path.Join("views", "general", "header.html")
	var filepath = path.Join("views", "main", "student.html")

	temp, err := template.New("student.html").Funcs(funcMap).ParseFS(t.embed, filepath, header)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	err = temp.Execute(c.Writer, dataTemplate)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
	}
}

func (t *studentWeb) StudentAddProcess(c *gin.Context) {
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
	// ID, _ := strconv.Atoi(c.Request.FormValue("id"))
	classID, _ := strconv.Atoi(c.Request.FormValue("class_id"))
	mahasiswa := model.Mahasiswa{
		// ID: ID,
		Name: c.Request.FormValue("name"),
		Address :c.Request.FormValue("address"),
		ClassId: classID,
	}

	status, err := t.studentClient.AddStudent(session.Token, mahasiswa)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message="+err.Error())
		return
	}

	if status == 201 {
		c.Redirect(http.StatusSeeOther, "/client/login")
	} else {
		c.Redirect(http.StatusSeeOther, "/client/modal?status=error&message=Add Student Failed!")
	}
}
