package versioning

import "github.com/gin-gonic/gin"

type VersionMap = map[string]gin.HandlerFunc

func (m VersionMap) getHandler(version string) (gin.HandlerFunc, bool) {
	handlerFunc, ok := m[version]

	return handlerFunc, ok
}
