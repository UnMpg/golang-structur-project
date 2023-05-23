package connect

import (
	"fmt"
	"gin-api-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DbPg *dbPgsql
)

type dbPgsql struct {
	dbpg *gorm.DB
}

func GetConnectionDB() (*gorm.DB, error) {
	return DbPg.GetConnection()
}

func (dbq *dbPgsql) GetConnection() (*gorm.DB, error) {
	return dbq.dbpg, nil
}

func InitConnectionDB() error {

	DbPg = new(dbPgsql)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.MyEnv.DBHost, config.MyEnv.DBUserPassword, config.MyEnv.DBUserPassword, config.MyEnv.DBName, config.MyEnv.DBPort)

	conn, errdb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errdb != nil {
		log.Fatal("Failed to connect database")
	}

	DbPg.dbpg = conn
	return nil
}

//func ConnectDB(c *config.Config) {
//	var err error
//	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUserPassword, c.DBUserPassword, c.DBName, c.DBPort)
//
//	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to connect database")
//	}
//
//	fmt.Println("Connect Successfully to the Database")
//
//}
