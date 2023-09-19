package logic

import (
	"context"
	"fmt"
	"strconv"

	"GopherTok/common/errorx"

	"github.com/pkg/errors"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowCountLogic {
	return &GetFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowCountLogic) GetFollowCount(in *pb.GetFollowCountReq) (*pb.GetFollowCountResp, error) {
	countCmd := l.svcCtx.Rdb.HGet(l.ctx, "followCount", fmt.Sprintf("%d:followCount", in.Userid))
	if countCmd.Err() != nil {
		if countCmd.Err().Error() != "redis: nil" {
			return nil,
				errors.Wrapf(errorx.NewDefaultError("redis get err:"+countCmd.Err().Error()), "redis get err ：%v", countCmd.Err())
		}

		if countCmd.Err().Error() == "redis: nil" {
			// 如果redis中没有则从mysql中拉取并更新至redis中
			db := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
				Where("follower_id = ?", in.Userid).Find(&[]pb.FollowSubject{})
			err := db.Error
			countMysql := db.RowsAffected

			if err != nil {
				return nil,
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			}

			l.svcCtx.Rdb.HSet(l.ctx, "followCount", fmt.Sprintf("%d:followCount", in.Userid), strconv.FormatInt(countMysql, 10))

			return &pb.GetFollowCountResp{
				StatusCode: 0,
				StatusMsg:  "get followCount successfully",
				Count:      countMysql,
			}, nil
		}
	}
	countInt, _ := strconv.ParseInt(countCmd.Val(), 10, 64)

	return &pb.GetFollowCountResp{
		StatusCode: 0,
		StatusMsg:  "get followCount successfully",
		Count:      countInt,
	}, nil
}
