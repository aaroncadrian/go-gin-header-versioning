package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VersionUtil struct {
	HeaderKey string
}

type VersionMap = map[string]gin.HandlerFunc

func (u VersionUtil) RegisterVersionMap(_ VersionMap) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		u.CheckIfHeaderIsPresent(),
	}
}

func (u VersionUtil) CheckIfHeaderIsPresent() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := u.HeaderKey

		if context.GetHeader(key) == "" {
			_ = context.AbortWithError(http.StatusBadRequest, fmt.Errorf(`"%v" header is required`, key))
		} else {
			context.Next()
		}
	}
}
