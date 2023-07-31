package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"GopherTok/server/user/api/internal/svc"
	"GopherTok/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	cnt, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.RegisterResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		UserId: cnt.UserId,
		Token:  cnt.Token,
	}, nil
}
