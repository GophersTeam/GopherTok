package logic

import (
	con "GopherTok/common/consts"
	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"
	"GopherTok/server/relation/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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
	to_userid, _ := strconv.ParseInt(req.UserId, 10, 64)

	rep, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &pb.GetFollowerReq{Userid: userid,
		ToUserId: to_userid})
	userlist := &[]types.User{}

	if err != nil {
		return nil, err
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
