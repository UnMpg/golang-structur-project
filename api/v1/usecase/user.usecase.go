package usecase

import (
	"fmt"
	"gin-api-test/api/v1/models"
	"gin-api-test/config"
	"gin-api-test/db/migrations"
	"gin-api-test/helpers/email"
	"gin-api-test/helpers/encode"
	"gin-api-test/helpers/jwt_token"
	"gin-api-test/helpers/log"
	"gin-api-test/helpers/password"
	"gin-api-test/repository/userRepository"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type UserUsecase struct {
	URepository userRepository.UserRepository
	DB          *gorm.DB
}

func NewUserUsecase(UserRepo userRepository.UserRepository) UserUsecase {
	return UserUsecase{URepository: UserRepo}
}

func (uc *UserUsecase) GetList(c *gin.Context) {
	fmt.Println("ini context", c)
	tes := uc.URepository.GetListUser(c, "INI dia adalaha isi")
	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "success", "success", tes))
}

func (uc *UserUsecase) SingUpUser(c *gin.Context) {
	var req *models.SignUp

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Log.Error("Error to binding request")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "fail", err.Error(), nil))
		return
	}

	if req.Password != req.PasswordConfirm {
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "fail", "Password not match", nil))
		return
	}

	hashedPassoword, err := password.HashPassword(req.Password)
	if err != nil {
		log.Log.Error("Error from hashedpassword")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "fail", err.Error(), nil))
		return
	}

	now := time.Now()
	newUser := migrations.User{
		Name:      req.Name,
		Email:     strings.ToLower(req.Email),
		Password:  hashedPassoword,
		Role:      "User",
		Verified:  true,
		Photo:     req.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := uc.URepository.SingUpUser(&newUser); err != nil {
		log.Log.Error("Error to Insert data")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "fail", err.Error(), nil))
		return
	}

	code := randstr.String(20)
	verificationCode := encode.Encode(code)

	newUser.VerificationCode = verificationCode
	uc.URepository.SingUpUserSaveVerification(newUser)

	var firstName = newUser.Name
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	emailData := email.EmailData{
		URL:       config.MyEnv.ClientOrigin + "verifyemail" + code,
		FirstName: firstName,
		Subject:   "Your Account Verification Code",
	}
	email.SendEmail(&newUser, &emailData)

	resultResponse := &migrations.User{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Photo:     newUser.Photo,
		Role:      newUser.Role,
		Provider:  newUser.Provider,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "success", "Data berhasil tersimpan", resultResponse))
}

func (uc *UserUsecase) SingInUser(c *gin.Context) {
	var req *models.SingIn

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Log.Error("Error binding JSON Input")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "Failed", err.Error(), nil))
		return
	}

	result, err := uc.URepository.GetUserByEmail(req.Email)
	if err != nil {
		log.Log.Error("User Not Found")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "failed", err.Error(), nil))
		return
	}

	if !result.Verified {
		log.Log.Error("Email not verify")
		c.JSON(http.StatusForbidden, models.CreateResponse(http.StatusForbidden, "failed", "Please verify your Email", nil))
		return
	}

	if err := password.VerifyPassword(result.Password, req.Password); err != nil {
		log.Log.Error("error password not match")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "fail", err.Error(), nil))
		return
	}

	accessToken, err := jwt_token.CreateToken(config.MyEnv.AccTokenExpireIn, result.ID, config.MyEnv.AccTokenPrivateKey)
	if err != nil {
		log.Log.Error("Error to create token")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "failed", err.Error(), nil))
		return
	}
	refreshToken, err := jwt_token.CreateToken(config.MyEnv.RefTokenExpireIn, result.ID, config.MyEnv.RefTokenPrivateKey)
	if err != nil {
		log.Log.Error("Error to create token")
		c.JSON(http.StatusBadRequest, models.CreateResponse(http.StatusBadRequest, "failed", err.Error(), nil))
		return
	}
	c.SetCookie("access_token", accessToken, config.MyEnv.AccTokenMaxEge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, config.MyEnv.RefTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.MyEnv.AccTokenMaxEge*60, "/", "localhost", false, false)
	respLogin := models.RespLogin{
		RoleUser: result.Role,
		Token:    accessToken,
	}
	c.JSON(http.StatusOK, models.CreateResponse(http.StatusOK, "Success", "success", respLogin))
}
