package handler

import (
	"net/http"

	"GopherTok/common/response"
	"GopherTok/server/user/api/internal/logic"
	"GopherTok/server/user/api/internal/svc"
	"GopherTok/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(r, w, resp, err) // ②

	}
}
