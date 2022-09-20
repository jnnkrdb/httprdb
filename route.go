package httprdb

import (
	"github.com/gin-gonic/gin"
)

// type struct implementation of the "route" type
//
// contains the request-type as string, the subpath, relative to
// the root-path as a string, and the function, which will be started,
// if the subpath gets called via the request
type Route struct {
	Request string
	SubPath string
	Handler gin.HandlerFunc
}
