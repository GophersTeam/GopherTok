package logic

import (
	"GopherTok/common/errorx"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendCountLogic {
	return &GetFriendCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendCountLogic) GetFriendCount(in *pb.GetFriendCountReq) (*pb.GetFriendCountResp, error) {
	count, err := l.svcCtx.Rdb.Hget("friendCount", fmt.Sprintf("%d:friendCount", in.Userid))
	if err != nil {
		return &pb.GetFriendCountResp{StatusCode: "-1",
				StatusMsg: err.Error(),
				Count:     0},
			errors.Wrapf(errorx.NewDefaultError("redis get err:"+err.Error()), "redis get err ：%v", err)
	}
	if count == "nil" {
		//如果redis中没有则从mysql中拉取并更新至redis中
		friend := []pb.FollowSubject{}
		err = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
			Where("user_id = ?", in.Userid).Find(&friend).Error
		if err != nil {
			return &pb.GetFriendCountResp{StatusCode: "-1",
					StatusMsg: err.Error(),
					Count:     0},
				errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
		}
		countMysql := 0
		for _, v := range friend {

			err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
				Where("user_id = ? AND follower_id = ?", v.FollowerId, in.Userid).First(&pb.FollowSubject{}).Error
			if err != nil {
				if err != gorm.ErrRecordNotFound {
					return &pb.GetFriendCountResp{StatusCode: "-1",
							StatusMsg: err.Error(),
							Count:     0},
						errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
				}
			}
			countMysql++
		}
		l.svcCtx.Rdb.Hset("friendCount", fmt.Sprintf("%d:friendCount", in.Userid), string(countMysql))

		return &pb.GetFriendCountResp{StatusCode: "0",
			StatusMsg: "get friendCount successfully",
			Count:     int64(countMysql)}, nil
	}
	countInt, err := strconv.ParseInt(count, 10, 64)

	return &pb.GetFriendCountResp{StatusCode: "0",
		StatusMsg: "get friendCount successfully",
		Count:     countInt}, nil
}
