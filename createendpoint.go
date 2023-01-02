package httprdb

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jnnkrdb/corerdb/prtcl"
)

// -----------------------------------------------------------

// booting an api-endpoint
//
// Parameters:
//   - `endpoint` : string > "[host]:[port]"
//   - `debug` : string > "release"/"debug"
//   - `routes` : []Route > the routes with their functions
func CreateApiEndpoint(endpoint, debug string, additionalHeaders []string, routes []Route) Endpoint {

	prtcl.Log.Println("creating an api-endpoint on ", endpoint)

	ep := Endpoint{
		Status: status{
			Code:    http.StatusProcessing, // 102
			Ready:   false,
			Healthy: false,
			Message: http.StatusText(http.StatusProcessing),
		},
	}

	// setting cors
	corsc := cors.DefaultConfig()

	corsc.AllowAllOrigins = true

	corsc.AllowBrowserExtensions = true

	corsc.AllowCredentials = true

	corsc.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

	corsc.AllowWebSockets = true

	corsc.AllowHeaders = additionalHeaders

	gin.DefaultWriter = prtcl.Log.Writer()

	gin.SetMode(debug)

	gin.DisableConsoleColor()

	ep.gin = gin.New()

	ep.gin.Use(cors.New(corsc))

	ep.gin.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		return fmt.Sprintf(
			"[RDB] %s Status: %d | Method: %s | Path: %s |  %d ms | response: %d Bytes\n",
			fmt.Sprintf(
				"%d/%02d/%02d %02d:%02d:%02d.%.6s",
				time.Now().Year(),
				time.Now().Month(),
				time.Now().Day(),
				time.Now().Hour(),
				time.Now().Minute(),
				time.Now().Second(),
				strconv.Itoa(time.Now().Nanosecond())),
			params.StatusCode,
			params.Method,
			params.Path,
			params.Latency,
			params.BodySize)
	}))

	ep.gin.Use(gin.Recovery())

	ep.httpserver = &http.Server{
		Addr:    endpoint,
		Handler: ep.gin,
	}

	// configure the subpaths
	for _, route := range routes {
		ep.gin.Handle(route.Request, route.SubPath, route.Handler)
	}

	return ep
}
