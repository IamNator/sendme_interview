package main

import (
	"net/http"
	"os"

	"github.com/IamNator/sendme_interview/internal/dao"
	"github.com/IamNator/sendme_interview/logger"
	"github.com/IamNator/sendme_interview/middleware"
	"github.com/IamNator/sendme_interview/router"
)

func main() {

	routes := router.Routes(dao.PostGres)
	port := os.Getenv("PORT")
	logger.Logger.Println("server running on port :" + port)

	http.ListenAndServe(":"+port, middleware.CorHandler(routes))
}
