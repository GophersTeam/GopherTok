package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"
	user "GopherTok/server/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *pb.GetFollowerReq) (*pb.GetFollowerResp, error) {
	follow := []pb.FollowSubject{}
	followerList := []pb.User{}
	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").Where("user_id = ?", in.ToUserId).Find(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return &pb.GetFollowerResp{
					StatusCode: "-1",
					StatusMsg:  err.Error(),
					UserList:   nil,
				},
				errors.Wrapf(errorx.NewDefaultError("mysql add err:"+err.Error()), "mysql add err ：%v", err)
		} else {
			return &pb.GetFollowerResp{
				StatusCode: "0",
				StatusMsg:  "get followerList successfully",
				UserList:   nil,
			}, nil
		}
	}

	for _, v := range follow {

		use, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:        v.FollowerId,
			CurrentId: in.Userid,
		})
		if err != nil {
			return &pb.GetFollowerResp{
					StatusCode: "-1",
					StatusMsg:  err.Error(),
					UserList:   nil,
				},
				errors.Wrapf(errorx.NewDefaultError("userInfo get err:"+err.Error()), "userInfo get err ：%v", err)
		}
		follower := pb.User{
			Id:              use.Id,
			Name:            use.Name,
			FollowCount:     use.FollowCount,
			FollowerCount:   use.FollowerCount,
			IsFollow:        use.IsFollow,
			Avatar:          use.Avatar,
			BackgroundImage: use.BackgroundImage,
			Signature:       use.Signature,
			TotalFavourited: use.TotalFavorited,
			WorkCount:       use.WorkCount,
			FavouriteCount:  use.FavoriteCount,
		}

		followerList = append(followerList, follower)
	}
	return &pb.GetFollowerResp{
		StatusCode: "0",
		StatusMsg:  "get followerList successfully",
		UserList:   &followerList,
	}, nil
}
