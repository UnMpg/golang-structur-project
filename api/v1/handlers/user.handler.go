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
	}

}

//func (Uh *UserHandler) SingUpUser(c *gin.Context) {
//	res := models.Response{StatusCode: http.StatusOK, Status: "success", Message: "Halaman Login aja ya"}
//	c.JSON(http.StatusOK, res)
//}

//func LoginUser(c *gin.Context) {
//	res := models.Response{StatusCode: http.StatusOK, Status: "success", Message: "Halaman Login"}
//	c.JSON(http.StatusOK, res)
//}
//
//func RegisterUser(c *gin.Context) {
//	var payload *models.SignUp
//	if err := c.ShouldBindJSON(&payload); err != nil {
//		res := models.Response{StatusCode: http.StatusBadRequest, Status: "fail", Message: "Error"}
//		c.JSON(http.StatusBadRequest, res)
//	}
//
//	if payload.Password != payload.PasswordConfirm {
//		res := models.Response{StatusCode: http.StatusBadRequest, Status: "fail", Message: "Password not match"}
//		c.JSON(http.StatusBadRequest, res)
//	}
//
//	res := models.Response{StatusCode: http.StatusOK, Status: "success", Message: "Halaman Register"}
//	c.JSON(http.StatusOK, res)
//}
