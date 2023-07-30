package handler

import (
	"GopherTok/common/response"
	"GopherTok/server/user/api/internal/logic"
	"GopherTok/server/user/api/internal/svc"
	"GopherTok/server/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo(&req)
		response.Response(r, w, resp, err) //②

	}
}
