package http

import (
	"github.com/my-open-falcon/agent/funcs"
	"net/http"
)

func configIoStatRoutes() {
	http.HandleFunc("/page/diskio", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, funcs.IOStatsForPage())
	})
}
