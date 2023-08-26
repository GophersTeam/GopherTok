package logic

import (
	con "GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListRes, err error) {
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
	//var currentId int64 = 1
	rep, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &pb.GetFollowerReq{Userid: req.UserId, CurrentId: currentId})
	userlist := []types.User{}

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	for _, val := range rep.UserList {
		usr := &types.User{}
		_ = copier.Copy(&usr, &val)

		userlist = append(userlist, *usr)
	}

	return &types.FollowerListRes{
		StatusMsg:  rep.StatusMsg,
		StatusCode: rep.StatusCode,
		UserList:   userlist,
	}, nil
}
