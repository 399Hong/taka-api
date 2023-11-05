package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"taka-api/internal/pkg/websocket"
	"taka-api/internal/svc"
)

func RegisterWsHandlers(server *rest.Server, serverCtx *svc.ServiceContext, hub *websocket.Hub) {
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			//if r.URL.Path != "/" {
			//	http.Error(w, "Not found", http.StatusNotFound)
			//	return
			//}
			//if r.Method != "GET" {
			//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			//	return
			//}
			logx.Error("serving html")
			http.ServeFile(w, r, "home.html")
		},
	})

	server.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodGet,
				Path:   "/ws",
				Handler: func(w http.ResponseWriter, r *http.Request) {
					websocket.ServeWs(hub, w, r)
				},
			},
		},
		//rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		//rest.WithPrefix("/v1"),
		//TODO not sure if we need timeout??
		//rest.WithTimeout(3000*time.Millisecond),
	)
}
