package logic

import (
	"context"

	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/user/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	// 先从redis查询是否有该userId
	exists, err := model.CheckUsernameExists(l.ctx, l.svcCtx.Rdb, in.Username)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "redis查询错误 err：%v", err)
	}
	if !exists {
		return nil, errors.Wrapf(errorx.NewDefaultError("用户名不存在，请先注册"), "用户名不存在，请先注册 LoginReq：%v", in)
	}

	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err = l.svcCtx.MysqlDb.Where("username = ?", in.Username).Find(&u).Error; err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "mysql查询错误 err：%v", err)
	}
	if u.Password != utils.Md5Password(in.Password, l.svcCtx.Config.Salt) {
		return nil, errors.Wrapf(errorx.NewDefaultError("密码错误"), "login密码错误错误 LoginReq：%v", in)
	}
	AccessToken, RefreshToken := utils.GetToken(u.Id, uuid.New().String(), l.svcCtx.Config.Token.AccessToken, l.svcCtx.Config.Token.RefreshToken)
	token := AccessToken + " " + RefreshToken
	return &user.LoginResp{
		UserId: u.Id,
		Token:  token,
	}, nil
}
