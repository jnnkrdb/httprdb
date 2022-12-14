package httprdb

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jnnkrdb/corerdb/prtcl"
)

// struct, which contains the http-service, statusvalues
// and the routes to the functions
type Endpoint struct {
	gin        *gin.Engine
	httpserver *http.Server
	Status     status
}

// boot the api endpoint
func (ep *Endpoint) Boot() {

	prtcl.Log.Println("starting api-endpoint:", ep.httpserver.Addr)

	// enable default functions
	ep.gin.Handle("GET", "/healthz", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, ep.Status)
	})

	ep.gin.Handle("GET", "/oof", func(ctx *gin.Context) {

		prtcl.Log.Println("Shutting down HTTP-Server")

		ep.Status.set(http.StatusServiceUnavailable, false, true)

		ctx.IndentedJSON(http.StatusProcessing, ep.Status)

		cctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		if err := ep.httpserver.Shutdown(cctx); err != nil {

			prtcl.PrintObject(ep, ctx, cctx, cancel, err)
		}
	})

	ep.Status.set(http.StatusOK, true, true)

	if err := ep.httpserver.ListenAndServe(); err != nil {

		if err != http.ErrServerClosed {

			prtcl.PrintObject(ep, err)
		}
	}

	ep.Status.set(http.StatusServiceUnavailable, false, false)

	prtcl.Log.Println("stopped api-endpoint:", ep.httpserver.Addr)
}
