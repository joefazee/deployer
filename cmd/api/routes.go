package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joefazee/autodeploy/pkg/domain"
)

func (srv *server) setupRouter() {

	if domain.IsDevelopment(srv.config.Environment) || domain.IsTesting(srv.config.Environment) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/health-check", srv.healthCheck)
	v1.GET("/posts", srv.listPosts)
	v1.GET("/users", srv.listUsers)
	v1.POST("/users", srv.createUser)

	if srv.asyncMux != nil {
		srv.registerAsyncHandlers()
	}

	srv.router = router
}
