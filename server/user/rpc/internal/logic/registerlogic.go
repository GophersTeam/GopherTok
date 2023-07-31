package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/user/model"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"

	"GopherTok/server/user/rpc/internal/svc"
	"GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	// 先从redis查询是否有该userId
	exists, err := l.checkUsernameExists(l.svcCtx.Rdb, in.Username)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "redis查询错误 err：%v", err)

	}

	if exists {
		logc.Info(l.ctx, "Username already exists.")
		return nil, errors.Wrapf(errorx.NewDefaultError("用户名已经存在，请更换用户名"), "用户名已经存在，请更换用户名 RegisterReq：%v", in)
	}
	err = l.registerUser(l.svcCtx.Rdb, in.Username)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "redis set 错误 err：%v", err)
	}

	u := model.User{
		Username: in.Username,
		// 加盐加密
		Password:        utils.Md5Password(in.Password, l.svcCtx.Config.Salt),
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
	}
	if err := l.svcCtx.MysqlDb.Create(&u).Error; err != nil {
		logx.Error(err)
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "gorm creat user 错误 err：%v", err)
	}
	// 生成token
	AccessToken, RefreshToken := utils.GetToken(u.ID, uuid.New().String(), l.svcCtx.Config.Token.AccessToken, l.svcCtx.Config.Token.RefreshToken)
	token := AccessToken + " " + RefreshToken
	return &user.RegisterResp{
		UserId: u.ID,
		Token:  token,
	}, nil
}

func (l *RegisterLogic) checkUsernameExists(client *redis.ClusterClient, username string) (bool, error) {
	// 查询 Redis 中是否存在该用户名
	exists, err := client.Exists(l.ctx, "username_"+username).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (l *RegisterLogic) registerUser(client *redis.ClusterClient, username string) error {
	// 存储用户名到 Redis 中
	_, err := client.Set(l.ctx, "username_"+username, "registered", 0).Result()
	return err
}
