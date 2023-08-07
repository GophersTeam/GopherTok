package logic

import (
	"GopherTok/server/relation/rpc/pb"
	"context"
	"fmt"

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
	//exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.UserId})
	//if err != nil {
	//	return &types.FollowRes{
	//		StatusCode: "-1",
	//		StatusMsg:  err.Error(),
	//	}, err
	//}
	//if exists.Exists == false {
	//	return &types.FollowRes{
	//		StatusCode: "-1",
	//		StatusMsg:  "user doesn't exist",
	//	}, nil
	//}

	rep, err := l.svcCtx.RelationRpc.GetFriendList(l.ctx, &pb.GetFriendListReq{Userid: req.UserId})
	userlist := &[]types.User{}

	if err != nil {
		if err != nil {
			fmt.Print(err)
			return &types.FriendListRes{
				StatusCode: "-1",
				StatusMsg:  err.Error(),
				UserList:   nil,
			}, err
		}
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

	return &types.FriendListRes{
		StatusMsg:  rep.StatusMsg,
		StatusCode: rep.StatusCode,
		UserList:   userlist,
	}, nil
}
