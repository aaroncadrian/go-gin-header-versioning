package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VersionUtil interface {
	RegisterVersionMap(versionMap VersionMap) []gin.HandlerFunc
}

type versionUtil struct {
	HeaderKey string
}

type VersionMap = map[string]gin.HandlerFunc

func NewVersionUtil(headerKey string) VersionUtil {
	return &versionUtil{
		HeaderKey: headerKey,
	}
}

func (u *versionUtil) RegisterVersionMap(_ VersionMap) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		u.checkIfHeaderIsPresent(),
	}
}

func (u *versionUtil) checkIfHeaderIsPresent() gin.HandlerFunc {
	return func(context *gin.Context) {
		key := u.HeaderKey

		if context.GetHeader(key) == "" {
			_ = context.AbortWithError(http.StatusBadRequest, fmt.Errorf(`"%v" header is required`, key))
		} else {
			context.Next()
		}
	}
}
