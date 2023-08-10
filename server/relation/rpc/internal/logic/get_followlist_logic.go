package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/user/rpc/types/user"
	"context"
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
	followList := []*pb.User{}
	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").Where("follower_id = ?", in.Userid).Find(&follow).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return &pb.GetFollowListResp{
					StatusCode: "-1",
					StatusMsg:  err.Error(),
					UserList:   nil,
				},
				errors.Wrapf(errorx.NewDefaultError("mysql add err:"+err.Error()), "mysql add err ：%v", err)
		} else {
			return &pb.GetFollowListResp{
				StatusCode: "0",
				StatusMsg:  "get followList successfully",
				UserList:   nil,
			}, nil
		}
	}

	for _, v := range follow {

		use, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:        v.UserId,
			CurrentId: in.Userid,
		})
		if err != nil {
			return &pb.GetFollowListResp{
					StatusCode: "-1",
					StatusMsg:  err.Error(),
					UserList:   nil,
				},
				errors.Wrapf(errorx.NewDefaultError("userInfo get err:"+err.Error()), "userInfo get err ：%v", err)
		}
		follow1 := pb.User{}

		_ = copier.Copy(&follow1, &use)

		followList = append(followList, &follow1)
	}
	return &pb.GetFollowListResp{
		StatusCode: "0",
		StatusMsg:  "get followList successfully",
		UserList:   followList,
	}, nil

}
