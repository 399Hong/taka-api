package company

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"taka-api/internal/logic/company"
	"taka-api/internal/svc"
	"taka-api/internal/types"
)

func RegisterCompanyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompanyRegisterationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := company.NewRegisterCompanyLogic(r.Context(), svcCtx)
		resp, err := l.RegisterCompany(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
