package logic

import (
	con "GopherTok/common/consts"
	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/rpc/types/user"
	"context"
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
	userid := l.ctx.Value(con.UserId).(int64)
	exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.UserId})
	if err != nil {
		return &types.FollowerListRes{
			StatusCode: "-1",
			StatusMsg:  err.Error(),
			UserList:   nil,
		}, err
	}
	if exists.Exists == false {
		return &types.FollowerListRes{
			StatusCode: "-1",
			StatusMsg:  "user doesn't exist",
			UserList:   nil,
		}, nil
	}

	exists, err = l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: userid})
	if err != nil {
		return &types.FollowerListRes{
			StatusCode: "-1",
			StatusMsg:  err.Error(),
			UserList:   nil,
		}, err
	}
	if exists.Exists == false {
		return &types.FollowerListRes{
			StatusCode: "-1",
			StatusMsg:  "user doesn't exist",
			UserList:   nil,
		}, nil
	}
	rep, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &pb.GetFollowerReq{Userid: userid,
		ToUserId: req.UserId})
	userlist := &[]types.User{}

	if err != nil {
		return &types.FollowerListRes{
			StatusCode: "-1",
			StatusMsg:  err.Error(),
			UserList:   nil,
		}, err
	}
	for _, val := range *rep.UserList {
		user := &types.User{
			Id:              val.Id,
			Name:            val.Name,
			FollowCount:     val.FollowCount,
			FollowerCount:   val.FollowerCount,
			IsFollow:        val.IsFollow,
			Avatar:          val.Avatar,
			BackgroundImage: val.BackgroundImage,
			Signature:       val.Signature,
			TotalFavourited: val.TotalFavourited,
			WorkCount:       val.WorkCount,
			FavouriteCount:  val.FavouriteCount,
		}
		*userlist = append(*userlist, *user)
	}

	return &types.FollowerListRes{
		StatusMsg:  rep.StatusMsg,
		StatusCode: rep.StatusCode,
		UserList:   userlist,
	}, nil
}
