package logic

import (
	"context"

	"GopherTok/common/errorx"
	"GopherTok/server/user/model"

	"github.com/pkg/errors"

	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCountLogic {
	return &AddCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCountLogic) AddCount(in *user.AddCountReq) (*user.AddCountResp, error) {
	// todo: add your logic here and delete this line
	result := l.svcCtx.MysqlDb.Model(&model.User{}).Where("id = ?", in.Id).Updates(map[string]interface{}{
		"follow_count":   in.FollowCount,
		"follower_count": in.FollowerCount,
		"friend_count":   in.FriendCount,
	})
	if result.Error != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql update err:"+result.Error.Error()), "mysql update err ：%v", result.Error)
	}
	return &user.AddCountResp{}, nil
}
