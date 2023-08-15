package handler

import (
	"GopherTok/common/response"
	"net/http"

	"GopherTok/server/relation/api/internal/logic"
	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FollowListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowListLogic(r.Context(), svcCtx)
		resp, err := l.FollowList(&req)
		response.Response(r, w, resp, err)
	}
}
