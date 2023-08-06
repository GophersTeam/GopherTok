package logic

import (
	"GopherTok/common/consts"
	"GopherTok/server/comment/api/internal/svc"
	"GopherTok/server/comment/api/internal/types"
	"GopherTok/server/comment/rpc/commentrpc"
	"context"
	"github.com/jinzhu/copier"

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
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.CommentListResponse)

	commentListResp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &commentrpc.GetCommentListRequest{
		VideoId: req.VideoId,
		UserId:  userId,
	})

	if err != nil {
		l.Errorf("Get comment list error: %v", err)
		return nil, err
	}
	resp.CommentList = make([]*types.Comment, len(commentListResp.CommentList))
	_ = copier.Copy(&resp.CommentList, &commentListResp.CommentList)
	return
}
