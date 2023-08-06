package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/user/model"

	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

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
	u := model.User{}
	isExists := false
	result := l.svcCtx.MysqlDb.Where("id = ?", in.Id).First(&u)
	// 判断记录是否存在
	if result.Error == gorm.ErrRecordNotFound {
		logc.Info(l.ctx, "ID %d 的记录不存在\n", in.Id)
	} else if result.Error != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("mysql查询出错，err:"+result.Error.Error()), "mysql查询出错，err:%v", result.Error)
	} else {
		logc.Info(l.ctx, "user_id %d 的记录存在\n", in.Id)
		isExists = true
	}
	return &user.UserIsExistsResp{
		Exists: isExists,
	}, nil
}
