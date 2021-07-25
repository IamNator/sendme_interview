package dao

import (
	"log"

	"github.com/IamNator/sendme_interview/config"
	"github.com/IamNator/sendme_interview/logger"
	"github.com/jinzhu/gorm"

	//	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

//PostGres ...
var PostGres *gorm.DB

//Connect creates a connection to PostGresl database
func Connect(conStr string) (*gorm.DB, error) {

	db, err := gorm.Open("postgres", conStr) // "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"

	if err != nil {
		return nil, err
	}

	logger.Logger.Println("database is open")

	return db, nil
}

func init() {
	dao, er := Connect(config.Config.PostgresConnectionURL) //connects to database server

	if er != nil {
		log.Fatal(er.Error() + " \nunable to connect to database")
		return
	}

	PostGres = dao

}
