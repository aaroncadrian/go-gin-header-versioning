package versioning

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Util interface {
	MapVersions(versionMap VersionMap) []gin.HandlerFunc
}

type headerVersionUtil struct {
	HeaderKey string
}

func NewHeaderVersioningUtil(headerKey string) Util {
	return &headerVersionUtil{
		HeaderKey: headerKey,
	}
}

func (u *headerVersionUtil) MapVersions(versions VersionMap) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		u.checkIfHeaderIsPresent(),
		u.handleVersion(versions),
	}
}

func (u headerVersionUtil) getVersion(context *gin.Context) string {
	return context.GetHeader(u.HeaderKey)
}

func (u headerVersionUtil) handleVersion(versionMap VersionMap) gin.HandlerFunc {
	return func(context *gin.Context) {
		version := u.getVersion(context)

		handler, ok := versionMap[version]

		if ok {
			handler(context)
		} else {
			_ = context.AbortWithError(http.StatusBadRequest, fmt.Errorf(`"%v" is not a supported version`, version))
		}
	}
}

func (u *headerVersionUtil) checkIfHeaderIsPresent() gin.HandlerFunc {
	return func(context *gin.Context) {
		if u.getVersion(context) == "" {
			_ = context.AbortWithError(http.StatusBadRequest, fmt.Errorf(`"%v" header is required`, u.HeaderKey))
		} else {
			context.Next()
		}
	}
}
