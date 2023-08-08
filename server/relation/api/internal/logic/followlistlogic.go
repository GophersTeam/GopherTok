package logic

import (
	con "GopherTok/common/consts"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"fmt"

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
	userid := l.ctx.Value(con.UserId).(int64)
	exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.UserId})
	if err != nil {
		return &types.FollowListRes{
			StatusCode: "-1",
			StatusMsg:  err.Error(),
			UserList:   nil,
		}, err
	}
	if exists.Exists == false {
		return &types.FollowListRes{
			StatusCode: "-1",
			StatusMsg:  "user doesn't exist",
			UserList:   nil,
		}, nil
	}

	exists, err = l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: userid})
	if err != nil {
		return &types.FollowListRes{
			StatusCode: "-1",
			StatusMsg:  err.Error(),
			UserList:   nil,
		}, err
	}
	if exists.Exists == false {
		return &types.FollowListRes{
			StatusCode: "-1",
			StatusMsg:  "user doesn't exist",
			UserList:   nil,
		}, nil
	}

	followList, err := l.svcCtx.RelationRpc.GetFollowList(l.ctx, &pb.GetFollowListReq{
		Userid:   userid,
		ToUserId: req.UserId,
	})
	if err != nil {
		fmt.Print(err)
		return &types.FollowListRes{
			StatusCode: "-1",
			StatusMsg:  err.Error(),
			UserList:   nil,
		}, err
	}
	userList := []types.User{}
	for _, v := range *followList.UserList {
		user := types.User{
			Id:              v.Id,
			Name:            v.Name,
			FollowCount:     v.FollowCount,
			FollowerCount:   v.FollowerCount,
			IsFollow:        v.IsFollow,
			Avatar:          v.Avatar,
			BackgroundImage: v.BackgroundImage,
			Signature:       v.Signature,
			TotalFavourited: v.TotalFavourited,
			WorkCount:       v.WorkCount,
			FavouriteCount:  v.FavouriteCount,
		}
		userList = append(userList, user)
	}
	return &types.FollowListRes{
		StatusCode: "0",
		StatusMsg:  "get followList successfully",
		UserList:   &userList,
	}, nil
}
