package logic

import (
	"context"
	"fmt"

	con "GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/rpc/types/user"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListReq) (resp *types.FollowListRes, err error) {
	currentId := l.ctx.Value(con.UserId).(int64)

	exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.UserId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if exists.Exists == false {
		return nil, errors.Wrapf(errorx.NewDefaultError("user doesn't exist"), "user doesn't exist%v", nil)
	}

	exists, err = l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: currentId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if exists.Exists == false {
		return nil, errors.Wrapf(errorx.NewDefaultError("user doesn't exist"), "user doesn't exist%v", nil)
	}

	// var currentId int64 = 1
	followList, err := l.svcCtx.RelationRpc.GetFollowList(l.ctx, &pb.GetFollowListReq{
		Userid:    req.UserId,
		CurrentId: currentId,
	})
	if err != nil {
		fmt.Print(err)
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	userList := []types.User{}
	for _, v := range followList.UserList {
		usr := types.User{}
		_ = copier.Copy(&usr, &v)
		userList = append(userList, usr)
	}
	return &types.FollowListRes{
		StatusCode: 0,
		StatusMsg:  "get followList successfully",
		UserList:   userList,
	}, nil
}
