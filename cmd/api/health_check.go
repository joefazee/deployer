package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (srv *server) healthCheck(ctx *gin.Context) {

	host, err := os.Hostname()

	if err != nil {
		host = "unknown"
	}

	healthOutput := struct {
		Status string `json:"status"`
		Env    string `json:"env"`
		Host   string `json:"host"`
	}{
		Status: "success",
		Env:    srv.config.Environment,
		Host:   host,
	}
	ctx.JSON(http.StatusOK, healthOutput)
}
