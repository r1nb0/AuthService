package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/r1nb0/UserService/api/controllers"
	"github.com/r1nb0/UserService/api/middleware"
	"github.com/r1nb0/UserService/api/validation"
	"github.com/r1nb0/UserService/configs"
	"github.com/r1nb0/UserService/infra"
	"github.com/r1nb0/UserService/pkg/logging"
	"github.com/r1nb0/UserService/pkg/metrics"
	"github.com/r1nb0/UserService/pkg/utils"
	"github.com/r1nb0/UserService/usecase"
	"net/http"
	"time"
)

type AppServer struct {
	httpServer     *http.Server
	cfg            *configs.Config
	logger         logging.Logger
	userController *controllers.UserController
	authMiddleware *middleware.AuthMiddleware
}

func NewAppServer(cfg *configs.Config) *AppServer {
	logger := logging.NewZapLogger(cfg)
	db, err := infra.InitPostgres(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	repo := infra.NewUserRepository(db, logger)
	jwtUtil := utils.NewJWTUtil(cfg)
	userUsecase := usecase.NewUserService(repo, jwtUtil, cfg)
	authMiddleware := middleware.NewAuthMiddleware(jwtUtil)
	userController := controllers.NewUserController(userUsecase)
	return &AppServer{
		cfg:            cfg,
		logger:         logger,
		authMiddleware: authMiddleware,
		userController: userController,
	}
}

func (serv *AppServer) initPrometheus() {
	err := prometheus.Register(metrics.DBCall)
	if err != nil {
		serv.logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		serv.logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
}

func (serv *AppServer) initRoutes(router *gin.Engine) {
	api := router.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/sign-in", serv.userController.SignIn)
		auth.POST("/sign-up", serv.userController.SignUp)
	}
	users := v1.Group("/users")
	{
		users.GET("/", serv.userController.GetAll)
		users.GET("/:id", serv.userController.GetByID)
		users.PUT("/", serv.authMiddleware.Authentication(), serv.userController.Update)
	}
	router.Use(middleware.PrometheusMiddleware())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// TODO impl
func (serv *AppServer) initSwagger() {

}

func (serv *AppServer) initValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		if err := val.RegisterValidation(
			"password",
			validation.PasswordValidator,
		); err != nil {
			serv.logger.Fatal(logging.Validation, logging.PasswordValidation, err.Error(), nil)
		}
	}
}

func (serv *AppServer) Run() {

	router := gin.Default()

	serv.initSwagger()
	serv.initValidators()
	serv.initPrometheus()
	serv.initRoutes(router)

	serv.httpServer = &http.Server{
		Addr:           ":" + serv.cfg.Server.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := serv.httpServer.ListenAndServe(); err != nil {
		serv.logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}
