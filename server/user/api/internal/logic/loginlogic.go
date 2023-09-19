package logic

import (
	"context"
	"fmt"

	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/user/rpc/types/user"

	"github.com/pkg/errors"

	"GopherTok/server/user/api/internal/svc"
	"GopherTok/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	cnt, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.LoginResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		UserId: cnt.UserId,
		Token:  cnt.Token,
	}, nil
}
