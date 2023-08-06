package logic

import (
	"GopherTok/server/comment/model"
	"GopherTok/server/comment/rpc/commentrpc"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"time"

	"GopherTok/server/comment/rpc/internal/svc"
	"GopherTok/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentRequest) (resp *pb.AddCommentResponse, err error) {
	// 对评论内容进行敏感词过滤
	in.Content = l.svcCtx.SensitiveWordFilter.Filter(in.Content)
	if in.Content == "" {
		return nil, errors.New("评论内容不能为空")
	}
	// 保存评论
	comment := &model.Comment{
		Id:         l.svcCtx.Snowflake.Generate().Int64(),
		Content:    in.Content,
		VideoId:    in.VideoId,
		UserId:     in.UserId,
		CreateDate: time.Now().Format(time.DateTime),
	}
	err = l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		l.Errorf("Insert comment error: %v", err)
		return
	}

	// 获取用户信息
	//userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoReq{
	//	UserId:       in.UserId,
	//	TargetUserId: in.UserId,
	//})
	//if err != nil {
	//	l.Errorf("Get user info error: %v", err)
	//	return
	//}
	resp = new(pb.AddCommentResponse)
	resp.Comment = new(pb.Comment)
	_ = copier.Copy(resp.Comment, comment)
	resp.Comment.User = new(commentrpc.User)
	//_ = copier.Copy(resp.Comment.User, userInfoResp.User)

	resp.Comment.User.Id = in.UserId
	resp.Comment.User.Username = "test"
	resp.Comment.User.Avatar = "https://static001.geekbang.org/account/avatar/00/19/61/0b/1c0b7f0d.jpg"

	return
}
