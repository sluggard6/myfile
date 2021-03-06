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
	"github.com/sluggard/myfile/store"
)

// HttpServer
type HttpServer struct {
	Config config.Config
	App    *iris.Application
	Store  store.Store
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
	var err error
	if s.Store, err = store.New(s.Config.Stroe.DataRoot); err != nil {
		log.Error(err.Error())
		return err
	}
	if err = model.Init(); err != nil {
		log.Error(err.Error())
		return err
	}
	s.RouteInit()
	return nil
}

// AuthRequired 登录验证
func AuthRequired(ctx iris.Context) {
	log.Debug(ctx.Request().Method, ctx.Request().RequestURI)
	//被忽略的url直接通过
	for _, v := range ignoreAuthUrl {
		if v == ctx.RequestPath(false) {
			ctx.Next()
		}
	}
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.Next()
}

// RouteInit 初始化路由
func (s *HttpServer) RouteInit() {

	app := s.App
	app.Options("/*", controller.Cors)
	// app.Party("/*", controller.Cors).AllowMethods(iris.MethodOptions)
	app.UseGlobal(controller.Cors)
	app.Use(AuthRequired, sess.Handler())
	mvc.New(app.Party("/test")).Handle(new(controller.TestController))
	mvc.New(app.Party("/user")).Handle(controller.NewUserController())
	mvc.New(app.Party("/library")).Handle(controller.NewLibraryController())
	mvc.New(app.Party("/folder")).Handle(controller.NewFolderController())
	mvc.New(app.Party("/file")).Handle(controller.NewFileController(s.Store))
	for _, route := range app.APIBuilder.GetRoutes() {
		log.Info(route)
	}
}
