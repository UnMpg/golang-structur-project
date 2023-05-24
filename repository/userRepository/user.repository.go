package userRepository

import (
	"fmt"
	"gin-api-test/db/migrations"
	"gin-api-test/helpers/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
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
	result := r.DB.Create(newUser)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate key Value uniq") {
		log.Log.Error("Error to Create Data")
		return result.Error
	} else if result.Error != nil {
		log.Log.Error("Error something bad happened")
		return result.Error
	}
	return nil
}

func (r *UserRepository) SingUpUserSaveVerification(newUser migrations.User) {
	r.DB.Save(newUser)
}

func (r *UserRepository) GetUserByEmail(email string) (migrations.User, error) {
	var user migrations.User
	record := r.DB.First(&user, "email = ?", strings.ToLower(email))
	if record.Error != nil {
		return user, record.Error
	}
	return user, nil
}
