package logic

import (
	"GopherTok/common/consts"
	"GopherTok/server/comment/api/internal/svc"
	"GopherTok/server/comment/api/internal/types"
	"GopherTok/server/comment/rpc/commentrpc"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionRequest) (resp *types.CommentActionResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.CommentActionResponse)
	if req.ActionType == consts.CommentAdd {
		addCommentResp, err := l.svcCtx.CommentRpc.AddComment(l.ctx, &commentrpc.AddCommentRequest{
			UserId:  userId,
			VideoId: req.VideoId,
			Content: req.CommentText,
		})
		if err != nil {
			return nil, err
		}
		resp.Comment = new(types.Comment)
		_ = copier.Copy(&resp.Comment, &addCommentResp.Comment)

	} else if req.ActionType == consts.CommentDel {
		if req.CommentId <= 0 {
			return nil, errors.New("id不合法")
		}

		_, err := l.svcCtx.CommentRpc.DelComment(l.ctx, &commentrpc.DelCommentRequest{
			CommentId: req.CommentId,
		})
		if err != nil {
			return nil, err
		}

	}
	return
}
