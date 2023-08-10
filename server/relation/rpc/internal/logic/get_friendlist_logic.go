package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *pb.GetFriendListReq) (*pb.GetFriendListResp, error) {
	friend := []pb.FollowSubject{}
	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ?", in.Userid).Find(&friend).Error
	if err != nil {
		return &pb.GetFriendListResp{StatusCode: "-1",
				StatusMsg: err.Error(),
				UserList:  nil},
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}
	friendList := []*pb.FriendUser{}
	for _, v := range friend {
		err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
			Where("user_id = ? AND follower_id = ?", v.FollowerId, in.Userid).First(&pb.FollowSubject{}).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return &pb.GetFriendListResp{StatusCode: "-1",
						StatusMsg: err.Error(),
						UserList:  nil},
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			} else {
				return &pb.GetFriendListResp{StatusCode: "0",
					StatusMsg: "get friendList successfully",
					UserList:  nil}, nil
			}

		}

		use, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:        v.FollowerId,
			CurrentId: in.Userid,
		})
		if err != nil {
			return &pb.GetFriendListResp{
					StatusCode: "-1",
					StatusMsg:  err.Error(),
					UserList:   nil,
				},
				errors.Wrapf(errorx.NewDefaultError("userInfo get err:"+err.Error()), "userInfo get err ：%v", err)
		}

		follow1 := pb.FriendUser{}
		_ = copier.Copy(&follow1, &use)
		friendList = append(friendList, &follow1)
	}

	return &pb.GetFriendListResp{StatusCode: "0",
		StatusMsg: "get friendList successfully",
		UserList:  friendList}, nil
}
