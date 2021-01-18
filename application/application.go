package application

import (
	stdContext "context"
	"fmt"
	"time"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	log "github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/application/controller"
	"github.com/sluggard/myfile/config"
	"github.com/sluggard/myfile/model"
)

// HttpServer
type HttpServer struct {
	Config config.Config
	App    *iris.Application
	Models []interface{}
	Status bool
}

var (
	gSessionId    = "GSESSIONID"
	sess          = sessions.New(sessions.Config{Cookie: gSessionId})
	ignoreAuthUrl = []string{
		"/test/ping",
		"/user/login",
		"/user/register",
	}
)

func initSwagger(app *iris.Application) {
	// url := swagger.URL("http://localhost:5678/swagger/doc.json") //The url pointing to API definition
	// app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))
	config := &swagger.Config{
		URL: "http://localhost:5678/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
}

func NewServer(config config.Config) *HttpServer {
	app := iris.New()
	initSwagger(app)
	// app.UseRouter(recover.New())
	// app.Use(AuthRequired())
	// app.Logger().SetLevel(libs.Config.LogLevel)
	// iris.RegisterOnInterrupt(func() {
	// sql, _ := easygorm.GetEasyGormDb().DB()
	// sql.Close()
	// })
	httpServer := &HttpServer{
		Config: config,
		App:    app,
		Status: false,
	}
	httpServer._Init()
	return httpServer
}

// Start
func (s *HttpServer) Start() error {
	if err := s.App.Run(
		// iris.Addr(fmt.Sprintf("%s:%d", libs.Config.Host, libs.Config.Port)),
		iris.Addr(fmt.Sprintf("%s:%d", s.Config.Server.Host, s.Config.Server.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(time.RFC3339),
	); err != nil {
		return err
	}
	s.Status = true
	return nil
}

// Start close the server at 3-6 seconds
func (s *HttpServer) Stop() {
	go func() {
		time.Sleep(3 * time.Second)
		ctx, cancel := stdContext.WithTimeout(stdContext.TODO(), 3*time.Second)
		defer cancel()
		s.App.Shutdown(ctx)
		s.Status = false
	}()
}

func (s *HttpServer) _Init() error {
	// err := libs.InitConfig(s.ConfigPath)
	// if err != nil {
	// 	logging.ErrorLogger.Errorf("系统配置初始化失败:", err)
	// 	return err
	// }
	// if libs.Config.Cache.Driver == "redis" {
	// 	cache.InitRedisCluster(libs.GetRedisUris(), libs.Config.Redis.Password)
	// }
	if err := model.Init(); err != nil {
		log.Error(err.Error())
		return err
	}
	s.RouteInit()
	return nil
}

func AuthRequired(ctx iris.Context) {
	log.Debug(ctx.Request().RequestURI)
	//被忽略的url直接通过
	for _, v := range ignoreAuthUrl {
		if v == ctx.RequestPath(false) {
			ctx.Next()
		}
	}
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	} else {
		ctx.Next()
	}
}

// RouteInit
func (s *HttpServer) RouteInit() {

	app := s.App

	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	// 	AllowCredentials: true,
	// 	AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// })
	app.Options("/*", controller.Cors)
	// app.Party("/*", controller.Cors).AllowMethods(iris.MethodOptions)
	app.UseGlobal(controller.Cors)
	app.Use(AuthRequired, sess.Handler())
	mvc.New(app.Party("/test")).Handle(new(controller.TestController))
	mvc.New(app.Party("/user")).Handle(controller.NewUserController())
	// log.Info(app.Macros().Lookup())

	// test.Handle(new(TestController))
	// test.Get("/ping", controller.GetPing)
	// test.Get("/help", help)
	// mvc.Configure(test, )

	// s.App.UseRouter(middleware.CrsAuth())
	// app := s.App.Party("/").AllowMethods(iris.MethodOptions)
	// {
	// 	app.HandleDir("/uploads", iris.Dir(filepath.Join(libs.CWD(), "uploads")))
	// 	v1 := app.Party("api/v1")
	// 	{
	// 		// 是否开启接口请求频率限制
	// 		if !libs.Config.Limit.Disable {
	// 			limitV1 := rate.Limit(libs.Config.Limit.Limit, libs.Config.Limit.Burst, rate.PurgeEvery(time.Minute, 5*time.Minute))
	// 			v1.Use(limitV1)
	// 		}
	// 		v1.Post("/admin/login", controllers.Login)
	// 		v1.PartyFunc("/admin", func(admin iris.Party) { //casbin for gorm                                                   // <- IMPORTANT, register the middleware.
	// 			admin.Use(middleware.JwtHandler().Serve, middleware.New().ServeHTTP) //登录验证
	// 			admin.Get("/logout", controllers.Logout).Name = "退出"
	// 			admin.Get("/expire", controllers.Expire).Name = "刷新 token"
	// 			admin.Get("/clear", controllers.Clear).Name = "清空 token"
	// 			admin.Get("/profile", controllers.Profile).Name = "个人信息"
	// 			admin.Post("/change_avatar", controllers.ChangeAvatar).Name = "修改头像"
	// 			admin.Post("/upload_file", iris.LimitRequestBodySize(libs.Config.MaxSize+1<<20), controllers.UploadFile).Name = "上传文件"

	// 			admin.PartyFunc("/users", func(users iris.Party) {
	// 				users.Get("/", controllers.GetUsers).Name = "用户列表"
	// 				users.Get("/{id:uint}", controllers.GetUser).Name = "用户详情"
	// 				users.Post("/", controllers.CreateUser).Name = "创建用户"
	// 				users.Post("/{id:uint}", controllers.UpdateUser).Name = "编辑用户"
	// 				users.Delete("/{id:uint}", controllers.DeleteUser).Name = "删除用户"
	// 			})
	// 			admin.PartyFunc("/roles", func(roles iris.Party) {
	// 				roles.Get("/", controllers.GetAllRoles).Name = "角色列表"
	// 				roles.Get("/{id:uint}", controllers.GetRole).Name = "角色详情"
	// 				roles.Post("/", controllers.CreateRole).Name = "创建角色"
	// 				roles.Post("/{id:uint}", controllers.UpdateRole).Name = "编辑角色"
	// 				roles.Delete("/{id:uint}", controllers.DeleteRole).Name = "删除角色"
	// 			})
	// 			admin.PartyFunc("/perms", func(permissions iris.Party) {
	// 				permissions.Get("/", controllers.GetAllPermissions).Name = "权限列表"
	// 				permissions.Get("/{id:uint}", controllers.GetPermission).Name = "权限详情"
	// 				permissions.Post("/", controllers.CreatePermission).Name = "创建权限"
	// 				permissions.Post("/{id:uint}", controllers.UpdatePermission).Name = "编辑权限"
	// 				permissions.Delete("/{id:uint}", controllers.DeletePermission).Name = "删除权限"
	// 			})
	//admin.PartyFunc("/configs", func(configs iris.Party) {
	//	configs.Get("/", controllers.GetAllConfigs).Name = "系统配置列表"
	//	configs.Get("/{key:string}", controllers.GetConfig).Name = "系统配置详情"
	//	configs.Post("/", controllers.CreateConfig).Name = "创建系统配置"
	//	configs.Post("/{id:uint}", controllers.UpdateConfig).Name = "编辑系统配置"
	//	configs.Delete("/{id:uint}", controllers.DeleteConfig).Name = "删除系统配置"
	//})
	// })
	// }
	// return nil
	// }
}
