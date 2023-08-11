package main

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/handler/api"
	"a21hc3NpZ25tZW50/handler/web"
	webserv "a21hc3NpZ25tZW50/handler/webserver"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"
	"embed"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "embed"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler     api.UserAPI
	CategoryAPIHandler api.CategoryAPI
	TaskAPIHandler     api.TaskAPI
	StudentAPIHandler api.StudentAPI
	ClassAPIHandler api.ClassAPI
	DosenAPIHandler api.DosenAPI
	MatkulAPIHandler api.MatkulAPI
}

type ClientHandler struct {
	AuthWeb      web.AuthWeb
	HomeWeb      web.HomeWeb
	DashboardWeb web.DashboardWeb
	TaskWeb      web.TaskWeb
	CategoryWeb  web.CategoryWeb
	StudentWeb	 web.StudentWeb
	ClassWeb	 web.ClassWeb
	DosenWeb 	 web.DosenWeb
	MatkulWeb	 web.MatkulWeb
	ServerWe 	 webserv.ServerWeb
	ModalWeb     web.ModalWeb
}

//go:embed views/*
var Resources embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode) //release

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()
		db := db.NewDB()
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] \"%s %s %s\"\n",
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())

		dbCredential := model.Credential{
			Host:         "localhost",
			Username:     "postgres",
			Password:     "rizwan123",
			DatabaseName: "school",
			Port:         5432,
			Schema:       "public",
		}

		conn, err := db.Connect(&dbCredential)
		if err != nil {
			panic(err)
		}

		conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Category{}, &model.Task{}, &model.Mahasiswa{}, &model.Class{}, &model.Dosen{}, &model.Matkul{})

		router = RunServer(conn, router)
		router = RunClient(conn, router, Resources)

		fmt.Println("Server is running on port 8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}

	}()

	wg.Wait()
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	userRepo := repo.NewUserRepo(db)
	sessionRepo := repo.NewSessionsRepo(db)
	categoryRepo := repo.NewCategoryRepo(db)
	taskRepo := repo.NewTaskRepo(db)
	studentRepo := repo.NewStudentRepo(db)
	classRepo := repo.NewClassRepo(db)
	dosenRepo := repo.NewDosenRepo(db)
	matkulRepo := repo.NewMatkulRepo(db)

	userService := service.NewUserService(userRepo, sessionRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	taskService := service.NewTaskService(taskRepo)
	studentService := service.NewStudentService(studentRepo)
	classService := service.NewClassService(classRepo)
	dosenService := service.NewDosenService(dosenRepo)
	matkulService := service.NewMatkulService(matkulRepo)


	userAPIHandler := api.NewUserAPI(userService)
	categoryAPIHandler := api.NewCategoryAPI(categoryService)
	taskAPIHandler := api.NewTaskAPI(taskService)
	studentAPIHandler := api.NewStudentAPI(studentService)
	classAPIHandler := api.NewClassAPI(classService)
	dosenAPIHandler := api.NewDosenAPI(dosenService)
	matkulAPIHandler := api.NewMatkulAPI(matkulService)

	apiHandler := APIHandler{
		UserAPIHandler:     userAPIHandler,
		CategoryAPIHandler: categoryAPIHandler,
		TaskAPIHandler:     taskAPIHandler,
		StudentAPIHandler: studentAPIHandler,
		ClassAPIHandler: classAPIHandler,
		DosenAPIHandler: dosenAPIHandler,
		MatkulAPIHandler: matkulAPIHandler,
	}

	version := gin.Group("/api/v1")
	{
		user := version.Group("/user")
		{
			user.POST("/login", apiHandler.UserAPIHandler.Login)
			user.POST("/register", apiHandler.UserAPIHandler.Register)

			user.Use(middleware.Auth())
			user.GET("/tasks", apiHandler.UserAPIHandler.GetUserTaskCategory)
		}

		task := version.Group("/task")
		{
			task.Use(middleware.Auth())
			task.POST("/add", apiHandler.TaskAPIHandler.AddTask)
			task.GET("/get/:id", apiHandler.TaskAPIHandler.GetTaskByID)
			task.PUT("/update/:id", apiHandler.TaskAPIHandler.UpdateTask)
			task.DELETE("/delete/:id", apiHandler.TaskAPIHandler.DeleteTask)
			task.GET("/list", apiHandler.TaskAPIHandler.GetTaskList)
			task.GET("/category/:id", apiHandler.TaskAPIHandler.GetTaskListByCategory)
		}

		category := version.Group("/category")
		{
			category.Use(middleware.Auth())
			category.POST("/add", apiHandler.CategoryAPIHandler.AddCategory)
			category.GET("/get/:id", apiHandler.CategoryAPIHandler.GetCategoryByID)
			category.PUT("/update/:id", apiHandler.CategoryAPIHandler.UpdateCategory)
			category.DELETE("/delete/:id", apiHandler.CategoryAPIHandler.DeleteCategory)
			category.GET("/list", apiHandler.CategoryAPIHandler.GetCategoryList)
		}

		student := version.Group("/student")
		{
			student.Use(middleware.Auth())
			student.POST("/add", apiHandler.StudentAPIHandler.AddStudent)
			student.GET("/get/:id", apiHandler.StudentAPIHandler.GetStudentByID)
			student.PUT("/update/:id", apiHandler.StudentAPIHandler.UpdateStudent)
			student.DELETE("/delete/:id", apiHandler.StudentAPIHandler.DeleteStudent)
			student.GET("/list", apiHandler.StudentAPIHandler.GetStudentList)
			student.GET("/class/:id", apiHandler.StudentAPIHandler.GetStudentListByClass)
		}

		class := version.Group("/class")
		{
			class.Use(middleware.Auth())
			class.POST("/add", apiHandler.ClassAPIHandler.AddClass)
			class.GET("/get/:id", apiHandler.ClassAPIHandler.GetClassByID)
			class.PUT("/update/:id", apiHandler.ClassAPIHandler.UpdateClass)
			class.DELETE("/delete/:id", apiHandler.ClassAPIHandler.DeleteClass)
			class.GET("/list", apiHandler.ClassAPIHandler.GetClassList)
		}
		dosen := version.Group("/dosen")
		{
			dosen.Use(middleware.Auth())
			dosen.POST("/add", apiHandler.DosenAPIHandler.AddDosen)
			dosen.GET("/get/:id", apiHandler.DosenAPIHandler.GetDosenByID)
			dosen.PUT("/update/:id", apiHandler.DosenAPIHandler.UpdateDosen)
			dosen.DELETE("/delete/:id", apiHandler.DosenAPIHandler.DeleteDosen)
			dosen.GET("/list", apiHandler.DosenAPIHandler.GetDosenList)
			dosen.GET("/matkul/:id", apiHandler.DosenAPIHandler.GetDosenListByMatkul)
		}
		matkul := version.Group("/matkul")
		{
			matkul.Use(middleware.Auth())
			matkul.POST("/add", apiHandler.MatkulAPIHandler.AddMatkul)
			matkul.GET("/get/:id", apiHandler.MatkulAPIHandler.GetMatkulByID)
			matkul.PUT("/update/:id", apiHandler.MatkulAPIHandler.UpdateMatkul)
			matkul.DELETE("/delete/:id", apiHandler.MatkulAPIHandler.DeleteMatkul)
			matkul.GET("/list", apiHandler.MatkulAPIHandler.GetMatkulList)
		}
		public := version.Group("/public")
		{
			public.Use(middleware.Auth())
			public.GET("/list", api.AmbilAPI)
		}

	}

	return gin
}

func RunClient(db *gorm.DB, engine *gin.Engine, embed embed.FS) *gin.Engine {
	sessionRepo := repo.NewSessionsRepo(db)
	sessionService := service.NewSessionService(sessionRepo)

	userClient := client.NewUserClient()
	taskClient := client.NewTaskClient()
	categoryClient := client.NewCategoryClient()
	studentClient := client.NewStudentClient()
	classClient := client.NewClassClient()
	dosenClient := client.NewDosenClient()
	matkulClient := client.NewMatkulClient()

	authWeb := web.NewAuthWeb(userClient, sessionService, embed)
	modalWeb := web.NewModalWeb(embed)
	homeWeb := web.NewHomeWeb(embed)
	dashboardWeb := web.NewDashboardWeb(userClient, sessionService, embed)
	taskWeb := web.NewTaskWeb(taskClient, sessionService, embed)
	categoryWeb := web.NewCategoryWeb(categoryClient, sessionService, embed)
	studentWeb := web.NewStudentWeb(studentClient, sessionService, embed)
	classWeb := web.NewClassWeb(classClient, sessionService, embed)
	dosenWeb := web.NewDosenWeb(dosenClient, sessionService, embed)
	serverWeb := webserv.NewServerWeb(embed)
	matkulWeb := web.NewMatkulWeb(matkulClient, sessionService, embed)

	client := ClientHandler{
		authWeb, homeWeb, dashboardWeb, taskWeb, categoryWeb, studentWeb, classWeb, dosenWeb, matkulWeb, serverWeb, modalWeb,
	}

	engine.StaticFS("/static", http.Dir("frontend/public"))

	engine.GET("/", client.HomeWeb.Index)

	user := engine.Group("/client")
	{
		user.GET("/login", client.AuthWeb.Login)
		user.POST("/login/process", client.AuthWeb.LoginProcess)
		user.GET("/register", client.AuthWeb.Register)
		user.POST("/register/process", client.AuthWeb.RegisterProcess)

		user.GET("/server", client.ServerWe.IndexServer)


		user.Use(middleware.Auth())
		user.GET("/logout", client.AuthWeb.Logout)

		
	}

	main := engine.Group("/client")
	{
		main.Use(middleware.Auth())
		main.GET("/dashboard", client.DashboardWeb.Dashboard)

		main.GET("/task", client.TaskWeb.TaskPage)
		user.POST("/task/add/process", client.TaskWeb.TaskAddProcess)
		user.PUT("/task/update", client.TaskWeb.TaskUpdateProcess)
		user.DELETE("/task/delete", client.TaskWeb.TaskDeleteProcess)

		main.GET("/category", client.CategoryWeb.Category)

		main.GET("/student", client.StudentWeb.StudentPage)
		user.POST("/student/add/process", client.StudentWeb.StudentAddProcess)

		main.GET("/class", client.ClassWeb.Class)

		main.GET("/dosen", client.DosenWeb.DosenPage)
		user.POST("/dosen/add/process", client.DosenWeb.DosenAddProcess)

		main.GET("/matkul", client.MatkulWeb.Matkul)
	}

	modal := engine.Group("/client")
	{
		modal.GET("/modal", client.ModalWeb.Modal)
	}

	return engine
}
