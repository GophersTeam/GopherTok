package handler

import (
	"GopherTok/common/response"
	"net/http"

	"GopherTok/server/relation/api/internal/logic"
	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.FollowerList(&req)
		response.Response(r, w, resp, err)
	}
}
