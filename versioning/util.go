package versioning

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Util interface {
	MapVersions(versionMap VersionMap) []gin.HandlerFunc
}

type versionUtil struct {
	HeaderKey string
}

func NewUtil(headerKey string) Util {
	return &versionUtil{
		HeaderKey: headerKey,
	}
}

func (u *versionUtil) MapVersions(versions VersionMap) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		u.checkIfHeaderIsPresent(),
		u.handleVersion(versions),
	}
}

func (u versionUtil) getVersion(context *gin.Context) string {
	return context.GetHeader(u.HeaderKey)
}

func (u versionUtil) handleVersion(versionMap VersionMap) gin.HandlerFunc {
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

func (u *versionUtil) checkIfHeaderIsPresent() gin.HandlerFunc {
	return func(context *gin.Context) {
		if u.getVersion(context) == "" {
			_ = context.AbortWithError(http.StatusBadRequest, fmt.Errorf(`"%v" header is required`, u.HeaderKey))
		} else {
			context.Next()
		}
	}
}
