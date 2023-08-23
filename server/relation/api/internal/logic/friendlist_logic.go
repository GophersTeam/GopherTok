package logic

import (
	con "GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListRes, err error) {
	userid := l.ctx.Value(con.UserId).(int64)
	exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.UserId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if exists.Exists == false {
		return nil, errors.Wrapf(errorx.NewDefaultError("user doesn't exist"), "user doesn't exist%v", nil)
	}
	//var userid int64 = 1
	rep, err := l.svcCtx.RelationRpc.GetFriendList(l.ctx, &pb.GetFriendListReq{Userid: userid})
	userlist := []types.FriendUser{}

	if err != nil {
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}
	}
	for _, val := range rep.UserList {
		usr := types.FriendUser{}
		_ = copier.Copy(&usr, &val)
		userlist = append(userlist, usr)
	}

	return &types.FriendListRes{
		StatusMsg:  rep.StatusMsg,
		StatusCode: rep.StatusCode,
		UserList:   userlist,
	}, nil
}
