package logic

import (
	"GopherTok/common/errorx"
	"context"
	"github.com/jinzhu/copier"
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
	var follow []pb.FollowSubject

	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").Where("user_id = ?", in.Userid).Find(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil,
				errors.Wrapf(errorx.NewDefaultError("mysql add err:"+err.Error()), "mysql add err ：%v", err)
		} else {
			return &pb.GetFollowerResp{
				StatusCode: 0,
				StatusMsg:  "get followerList successfully",
				UserList:   nil,
			}, nil
		}
	}

	var (
		followerList []*pb.User
		followerChan = make(chan *pb.User, len(follow))
		errChan      = make(chan error, len(follow))
	)
	for i := 0; i < len(follow); i++ {
		go func(i int) {

			use, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
				Id:        follow[i].FollowerId,
				CurrentId: in.Userid,
			})
			if err != nil {
				errChan <- errors.Wrapf(errorx.NewDefaultError("userInfo get err:"+err.Error()), "userInfo get err ：%v", err)
			}
			follower := &pb.User{}
			_ = copier.Copy(follower, &use)
			followerChan <- follower
		}(i)
	}
	for i := 0; i < len(follow); i++ {
		select {
		case follower := <-followerChan:
			followerList = append(followerList, follower)
		case err := <-errChan:
			return nil, err
		}

	}
	return &pb.GetFollowerResp{
		StatusCode: 0,
		StatusMsg:  "get followerList successfully",
		UserList:   followerList,
	}, nil
}
