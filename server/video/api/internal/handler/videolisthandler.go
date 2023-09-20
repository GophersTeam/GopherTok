package handler

import (
	"net/http"

	"GopherTok/common/response"
	"GopherTok/server/video/api/internal/logic"
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVideoListLogic(r.Context(), svcCtx)
		resp, err := l.VideoList(&req)
		response.Response(r, w, resp, err) // ②
	}
}
