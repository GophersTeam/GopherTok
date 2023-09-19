package logic

import (
	"GopherTok/common/errorx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

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
	isExists := false
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err == sqlx.ErrNotFound {
		logc.Info(l.ctx, "ID %d 的记录不存在\n", in.Id)
	} else if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询出错，err:"+err.Error()), "mysql查询出错，err:%v", err)
	} else {
		logc.Info(l.ctx, "user_id %d 的记录存在\n", in.Id)
		isExists = true
	}
	return &user.UserIsExistsResp{
		Exists: isExists,
	}, nil
}
