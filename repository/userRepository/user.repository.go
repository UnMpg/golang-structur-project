package userRepository

import (
	"fmt"
	"gin-api-test/api/v1/models"
	"gin-api-test/db/migrations"
	"gin-api-test/helpers/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (r *UserRepository) GetListUser(c *gin.Context, a string) string {
	coba := a
	fmt.Println("List user ini lahnya")
	return coba
}

func (r *UserRepository) SingUpUser(newUser *migrations.User) error {
	fmt.Println("report usernya", &newUser)
	result := r.DB.Create(&newUser)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate key Value uniq") {
		log.Log.Error("Error to Create Data")
		return result.Error
	} else if result.Error != nil {
		log.Log.Error("Error something bad happened")
		return result.Error
	}
	fmt.Println("hasil report singUp user", result)
	return nil
}

func (r *UserRepository) SingUpUserSaveVerification(newUser migrations.User) error {
	save := r.DB.Save(newUser)
	if save.Error != nil {
		log.Log.Error("Error to save gorm", save.Error)
		return save.Error
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (migrations.User, error) {
	var user migrations.User
	record := r.DB.First(&user, "email = ?", strings.ToLower(email))
	if record.Error != nil {
		return user, record.Error
	}
	return user, nil
}

func (r *UserRepository) GetUserList() ([]migrations.User, error) {
	var user []migrations.User

	record := r.DB.Find(&user, "role = ?", "user")
	if record.Error != nil {
		return nil, record.Error
	}
	return user, nil

}

func (r *UserRepository) UpdateUser(uid string, updateUser models.UpdateUser) (models.UpdateUser, error) {
	var user migrations.User

	record := r.DB.First(&user, "id = ?", uid)
	if record.Error != nil {
		return updateUser, record.Error
	}
	now := time.Now()
	updateUserSet := models.UpdateUser{
		Name:      updateUser.Name,
		Email:     updateUser.Email,
		Photo:     updateUser.Photo,
		UpdatedAt: now,
	}
	r.DB.Model(&user).Updates(updateUserSet)

	return updateUserSet, nil
}

func (r *UserRepository) DeleteUser(userId string) error {
	deleteUSer := r.DB.Delete(&migrations.User{}, "id = ?", userId)
	if deleteUSer.Error != nil {
		log.Log.Error("Error gorm delete User", deleteUSer.Error)
		return deleteUSer.Error
	}

	return nil
}
