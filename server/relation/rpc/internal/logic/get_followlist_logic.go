package logic

import (
	"context"

	"GopherTok/common/errorx"
	"GopherTok/server/user/rpc/types/user"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *pb.GetFollowListReq) (*pb.GetFollowListResp, error) {
	follow := []pb.FollowSubject{}
	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").Where("follower_id = ?", in.Userid).Find(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil,
				errors.Wrapf(errorx.NewDefaultError("mysql add err:"+err.Error()), "mysql add err ：%v", err)
		} else {
			return &pb.GetFollowListResp{
				StatusCode: 0,
				StatusMsg:  "get followList successfully",
				UserList:   nil,
			}, nil
		}
	}
	var (
		followList []*pb.User
		followChan = make(chan *pb.User, len(follow))
		errChan    = make(chan error, len(follow))
	)
	for i := 0; i < len(follow); i++ {
		go func(i int) {
			use, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
				Id:        follow[i].UserId,
				CurrentId: in.CurrentId,
			})
			if err != nil {
				errChan <- errors.Wrapf(errorx.NewDefaultError("userInfo get err:"+err.Error()), "userInfo get err ：%v", err)
			}
			follow1 := &pb.User{}

			_ = copier.Copy(follow1, &use)
			followChan <- follow1
		}(i)
	}
	for i := 0; i < len(follow); i++ {
		select {
		case follow := <-followChan:
			followList = append(followList, follow)
		case err := <-errChan:
			return nil, err
		}
	}

	return &pb.GetFollowListResp{
		StatusCode: 0,
		StatusMsg:  "get followList successfully",
		UserList:   followList,
	}, nil
}
