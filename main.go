package main

import (
	"net/http"

	"github.com/IamNator/sendme_interview/config"
	"github.com/IamNator/sendme_interview/internal/dao"
	"github.com/IamNator/sendme_interview/logger"
	"github.com/IamNator/sendme_interview/middleware"
	"github.com/IamNator/sendme_interview/router"
)

func main() {

	routes := router.Routes(dao.PostGres)
	port := config.Config.PORT
	logger.Logger.Println("server running on port :" + port)

	http.ListenAndServe(":"+port, middleware.CorHandler(routes))
}
