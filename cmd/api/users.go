package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joefazee/autodeploy/pkg/domain"
	"github.com/joefazee/autodeploy/pkg/tasks"
)

func (srv *server) listUsers(ctx *gin.Context) {

	users, err := srv.store.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  users,
	})
}

func (srv *server) createUser(ctx *gin.Context) {

	var req domain.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := srv.store.CreateUser(domain.User{
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = srv.sendJob(tasks.TaskUserCreated, u, nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  u,
	})
}

func (srv *server) longJob() {
	time.Sleep(10 * time.Second)
}
