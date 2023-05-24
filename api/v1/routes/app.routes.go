package routes

import (
	"gin-api-test/api/v1/handlers"
	"gin-api-test/api/v1/middleware"
	"gin-api-test/api/v1/usecase"
	"gin-api-test/db/connect"
	"gin-api-test/helpers/log"
	"gin-api-test/repository/userRepository"
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func CreateRoute(isDev bool) *gin.Engine {
	router := gin.New()
	router.Use(requestid.New())
	router.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		middleware.HandlePanic(c, err)
	}))

	router.Use(healthcheck.Default())
	router.Use(log.RequestLoggerActivity())
	return router
}

func RouteV1(r *gin.Engine) {
	dbPg, err := connect.GetConnectionDB()
	if err != nil {
		log.Log.Error("Error Connection to Postgres ", err.Error())
	}

	v1Userroute := r.Group("/api/v1")

	UserRepository := userRepository.NewUserRepository(dbPg)
	UserUsecase := usecase.NewUserUsecase(UserRepository)
	handlers.NewUserHandlerV1(v1Userroute, UserUsecase)
	//v1Userroute.GET("/", func(c *gin.Context) {
	//	log.Log.Info("Nama saya", "ini aja ", "tes dulu")
	//	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "success", "Halaman User", nil))
	//
	//})

}

func PrivateV1(r *gin.Engine) {
	dbPg, err := connect.GetConnectionDB()
	if err != nil {
		log.Log.Error("Error Connection to Postgres ", err.Error())
	}
	v1UserPrivate := r.Group("/private/v1")
	v1UserPrivate.Use(middleware.CekUserMiddleware())

	UserRepository := userRepository.NewUserRepository(dbPg)
	UserUsecase := usecase.NewUserUsecase(UserRepository)
	handlers.NewPrivateHandlerV1(v1UserPrivate, UserUsecase)

}
