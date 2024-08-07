//go:generate statik -src=./assets/dist
//go:generate go fmt statik/statik.go

package application

import (
	stdContext "context"
	"fmt"
	"strings"
	"time"

	"github.com/sluggard/myfile/service"
	//_ "github.com/sluggard/myfile/statik" // TODO: Replace with the absolute import path

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/rakyll/statik/fs"
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
		"/app",
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

func TokenRequired(ctx iris.Context) {
	path := config.GetConfig().Server.ContextPath
	//被忽略的url直接通过
	for _, v := range ignoreAuthUrl {
		if strings.HasPrefix(ctx.RequestPath(false), path+v) {
			// if path+v == ctx.RequestPath(false) {
			ctx.Next()
			return
		}
	}
	token := ctx.GetHeader("Authorization")
	if user, err := service.NewTokenService().CheckToken(token); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"code":    401,
			"message": "token验证失败",
		})
	} else {
		ctx.Values().Set("user", user)
		ctx.Next()
	}
}

// AuthRequired 登录验证
func AuthRequired(ctx iris.Context) {
	// iris.CookieSameSite = iris.SameSiteNoneMode
	log.Debug(ctx.Request().Method, "  ", ctx.Request().RequestURI)
	// log.Debug(ctx.RequestPath(false))
	session := sess.Start(ctx, iris.CookieSameSite(iris.SameSiteNoneMode))
	// log.Debug(session)
	// log.Debug(sess.GetCookieOptions())
	path := config.GetConfig().Server.ContextPath
	//被忽略的url直接通过
	for _, v := range ignoreAuthUrl {
		if strings.HasPrefix(ctx.RequestPath(false), path+v) {
			// if path+v == ctx.RequestPath(false) {
			ctx.Next()
			return
		}
	}
	if auth, _ := session.GetBoolean("authenticated"); !auth {
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
	app.Use(TokenRequired)
	// app.Use(AuthRequired, sess.Handler())
	// app.Use(sess.Handler())
	statikFS, err := fs.New()
	if err == nil {
		app.HandleDir(s.Config.Server.ContextPath+"/app", statikFS)
	} else {
		fmt.Printf("err: %v\n", err)
	}
	mvc.New(app.Party(s.Config.Server.ContextPath + "/test")).Handle(new(controller.TestController))
	mvc.New(app.Party(s.Config.Server.ContextPath + "/user")).Handle(controller.NewUserController())
	mvc.New(app.Party(s.Config.Server.ContextPath + "/library")).Handle(controller.NewLibraryController())
	mvc.New(app.Party(s.Config.Server.ContextPath + "/folder")).Handle(controller.NewFolderController())
	mvc.New(app.Party(s.Config.Server.ContextPath + "/file")).Handle(controller.NewFileController(s.Store))
	for _, route := range app.APIBuilder.GetRoutes() {
		log.Info(route)
	}
}

// func (s *HttpServer) addWebDAV(path string, libraryId int) {
// 	s.App.HandleDir()
// }
