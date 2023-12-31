// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"taka-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/signup",
				Handler: SignUpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/Login",
				Handler: LoginHandler(serverCtx),
			},
		},
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/profile/elser",
				Handler: GetProfileHandler(serverCtx),
			},
		},
		rest.WithTimeout(3000*time.Millisecond),
	)
}
