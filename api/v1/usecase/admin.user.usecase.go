package usecase

import (
	"fmt"
	"gin-api-test/api/v1/models"
	"gin-api-test/helpers/jwt_token"
	"gin-api-test/helpers/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *UserUsecase) GetProfileUser(c *gin.Context) {
	fmt.Println("kontestnya", c)
	_, err := jwt_token.GetUserID(c)
	if err != nil {
		log.Log.Error("Error tu get User Id", err)
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "Not Authorizator", err.Error(), nil))
		return
	}

	getUserList, err := uc.URepository.GetUserList()
	if err != nil {
		log.Log.Error("Error to get User List ::", err.Error())
		c.JSON(http.StatusInternalServerError, models.CreateResponse(http.StatusInternalServerError, "faile", err.Error(), nil))
		return
	}
	fmt.Println(getUserList)
	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "success", "success", getUserList))

}

func (uc *UserUsecase) UpdateUser(c *gin.Context) {
	var req models.UpdateUser
	uid, err := jwt_token.GetUserID(c)
	if err != nil {
		log.Log.Error("Error tu get User Id", err)
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "Not Authorizator", err.Error(), nil))
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Log.Error("Error tu Binding json data Id", err)
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "Not Authorizator", err.Error(), nil))
		return
	}

	updateUser, err := uc.URepository.UpdateUser(uid, req)
	if err != nil {
		log.Log.Error("Error to Update User List ::", err.Error())
		c.JSON(http.StatusInternalServerError, models.CreateResponse(http.StatusInternalServerError, "faile", err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "success", "success", updateUser))

}

func (uc *UserUsecase) DeleteUser(c *gin.Context) {
	_, err := jwt_token.GetUserID(c)
	if err != nil {
		log.Log.Error("Error tu get User Id", err)
		c.JSON(http.StatusUnauthorized, models.CreateResponse(http.StatusUnauthorized, "Not Authorizator", err.Error(), nil))
		return
	}

	deleteUserId := c.Param("userId")

	if err := uc.URepository.DeleteUser(deleteUserId); err != nil {
		log.Log.Error("Error tu get User Id", err)
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "failed", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "success", "success delete User By ID", deleteUserId))
}
