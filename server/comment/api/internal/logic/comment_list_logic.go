package logic

import (
	"GopherTok/server/comment/rpc/commentrpc"
	"context"
	"github.com/jinzhu/copier"

	"GopherTok/server/comment/api/internal/svc"
	"GopherTok/server/comment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	resp = new(types.CommentListResponse)
	commentListResp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &commentrpc.GetCommentListRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	resp.CommentList = make([]*types.Comment, 0, len(commentListResp.CommentList))
	copier.Copy(resp.CommentList, commentListResp.CommentList)

	return
}
