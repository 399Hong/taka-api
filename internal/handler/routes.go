// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	company "taka-api/internal/handler/company"
	user "taka-api/internal/handler/user"
	"taka-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/signup",
				Handler: user.SignUpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
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

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/company/register",
				Handler: company.RegisterCompanyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)
}
