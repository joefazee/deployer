package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *server) listPosts(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"posts": []string{
			"test", "test23", "test post", "more",
		},
	})
}
