package handlers

import (
	"gin-api-test/api/v1/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Uusecase usecase.UserUsecase
}

func NewUserHandlerV1(g *gin.RouterGroup, ru usecase.UserUsecase) {
	//handler := &UserHandler{Uusecase: ru}

	user := g.Group("/user")
	{
		user.GET("/", ru.GetList)
		user.POST("/register", ru.SingUpUser)
		user.POST("/login", ru.SingInUser)

	}

}

func NewPrivateHandlerV1(g *gin.RouterGroup, ru usecase.UserUsecase) {
	admin := g.Group("/admin")
	{
		admin.GET("/", ru.GetList)
		admin.GET("/list-user", ru.GetProfileUser)
		admin.PUT("/update-user", ru.UpdateUser)
		admin.DELETE("/delete-user/:userId", ru.DeleteUser)
	}
}
