package handler

import (
	"net/http"

	"GopherTok/common/response"

	"GopherTok/server/comment/api/internal/logic"
	"GopherTok/server/comment/api/internal/svc"
	"GopherTok/server/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		response.Response(r, w, resp, err)
	}
}
