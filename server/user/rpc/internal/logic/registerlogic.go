package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/user/model"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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
	exists, err := model.CheckUsernameExists(l.ctx, l.svcCtx.Rdb, in.Username)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "redis查询错误 err：%v", err)

	}
	if exists {
		logc.Info(l.ctx, "Username already exists.")
		return nil, errors.Wrapf(errorx.NewDefaultError("用户名已经存在，请更换用户名"), "用户名已经存在，请更换用户名 RegisterReq：%v", in)
	}
	err = model.RegisterUser(l.ctx, l.svcCtx.Rdb, in.Username)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "redis set 错误 err：%v", err)
	}

	u := model.User{
		Id:       l.svcCtx.Snowflake.Generate().Int64(),
		Username: in.Username,
		// 加盐加密
		Password: utils.Md5Password(in.Password, l.svcCtx.Config.Salt),
	}
	// kafka异步处理msg
	uMsg, err := json.Marshal(&u)
	if err != nil {
		logx.Errorf("json.Marshal msgs: %v error: %v", uMsg, err)
		return nil, errors.Wrapf(errorx.NewDefaultError("userinfo jons转换错误"), "用户名已经存在，请更换用户名 RegisterReq：%v", in)

	}
	if err = l.svcCtx.KqPusherClient.Push(string(uMsg)); err != nil {
		logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(uMsg), err)
		return nil, errors.Wrapf(errorx.NewDefaultError("userinfo 写入kafka错误"), "用户名已经存在，请更换用户名 RegisterReq：%v", in)

	}
	//if err := l.svcCtx.MysqlDb.Create(&u).Error; err != nil {
	//	logx.Error(err)
	//	return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "gorm creat user 错误 err：%v", err)
	//}
	// 生成token
	AccessToken, RefreshToken := utils.GetToken(u.Id, uuid.New().String(), l.svcCtx.Config.Token.AccessToken, l.svcCtx.Config.Token.RefreshToken)
	token := AccessToken + " " + RefreshToken
	return &user.RegisterResp{
		UserId: u.Id,
		Token:  token,
	}, nil
}
