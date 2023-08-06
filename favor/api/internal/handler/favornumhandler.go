package handler

import (
	"net/http"

	"GopherTok/server/favor/api/internal/logic"
	"GopherTok/server/favor/api/internal/svc"
	"GopherTok/server/favor/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavorNumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavorNumReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavorNumLogic(r.Context(), svcCtx)
		resp, err := l.FavorNum(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
