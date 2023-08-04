package logic

import (
	"GopherTok/common/consts"
	"GopherTok/server/comment/rpc/commentrpc"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"GopherTok/server/comment/api/internal/svc"
	"GopherTok/server/comment/api/internal/types"

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
	resp.Comment = new(types.Comment)
	if req.ActionType == consts.CommentAdd {
		// 对评论内容进行过滤
		req.CommentText = l.filterComment(req.CommentText)
		addCommentResp, err := l.svcCtx.CommentRpc.AddComment(l.ctx, &commentrpc.AddCommentRequest{
			UserId:  userId,
			VideoId: req.VideoId,
			Content: req.CommentText,
		})
		if err != nil {
			return nil, err
		}

		_ = copier.Copy(resp.Comment, addCommentResp.Comment)
		// 获取用户信息
		//userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
		//	UserId:       userId,
		//	TargetUserId: userId,
		//})
		//if err != nil {
		//	return nil, err
		//}
		//
		//_ = copier.Copy(resp.Comment.User, userInfoResp.User)
		resp.Comment.User = &types.User{
			Id:              userId,
			Username:        "aa",
			Avatar:          "bb",
			FollowCount:     123,
			TotalFavorited:  456,
			Signature:       "789",
			BackgroundImage: "111",
			FollowerCount:   222,
			WorkCount:       333,
			FavoriteCount:   444,
			IsFollow:        true,
		}

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

func (l *CommentActionLogic) filterComment(text string) string {
	return text
}
