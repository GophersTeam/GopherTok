package logic

import (
	"GopherTok/server/relation/rpc/pb"
	"context"
	"fmt"

	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowRes, err error) {
	//userid := l.ctx.Value(con.UserId).(int64)
	//exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.ToUserId})
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
	//
	//exists, err = l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: userid})
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
	var userid int64 = 1
	isFollow, err := l.svcCtx.RelationRpc.CheckIsFollow(l.ctx, &pb.CheckIsFollowReq{
		UserId:   userid,
		ToUserId: req.ToUserId,
	})
	if err != nil {
		fmt.Print(err)
		return &types.FollowRes{
			StatusCode: isFollow.StatusCode,
			StatusMsg:  isFollow.StatusMsg,
		}, err
	}

	if isFollow.IsFollow {
		return &types.FollowRes{
			StatusCode: "-1",
			StatusMsg:  "follow has exist",
		}, nil
	}

	follow, err := l.svcCtx.RelationRpc.AddFollow(l.ctx, &pb.AddFollowReq{
		UserId:   userid,
		ToUserId: req.ToUserId,
	})
	if err != nil {
		fmt.Print(err)
		return &types.FollowRes{
			StatusCode: follow.StatusCode,
			StatusMsg:  follow.StatusMsg,
		}, err
	}

	return &types.FollowRes{
		StatusCode: follow.StatusCode,
		StatusMsg:  follow.StatusMsg,
	}, nil
}
