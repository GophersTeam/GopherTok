package logic

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"GopherTok/server/user/api/internal/svc"
	"GopherTok/server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	// 获取userId
	userId, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	fmt.Println("user id =", userId)
	userCnt, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.UserInfoResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		ID:              userCnt.Id,
		Name:            userCnt.Name,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: userCnt.BackgroundImage,
		Signature:       userCnt.Signature,
		TotalFavorited:  "",
		WorkCount:       0,
		FavoriteCount:   0,
	}, nil
}
