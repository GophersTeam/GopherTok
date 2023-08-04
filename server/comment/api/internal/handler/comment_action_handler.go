package handler

import (
	"net/http"

	"GopherTok/server/comment/api/internal/logic"
	"GopherTok/server/comment/api/internal/svc"
	"GopherTok/server/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentActionLogic(r.Context(), svcCtx)
		resp, err := l.CommentAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
