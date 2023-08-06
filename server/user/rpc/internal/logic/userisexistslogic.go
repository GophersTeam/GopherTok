package logic

import (
	"context"

	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIsExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIsExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIsExistsLogic {
	return &UserIsExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserIsExistsLogic) UserIsExists(in *user.UserIsExistsReq) (*user.UserIsExistsResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserIsExistsResp{}, nil
}
