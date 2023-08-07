package logic

import (
	"GopherTok/common/errorx"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strconv"

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
	count, err := l.svcCtx.Rdb.Hget("followCount", fmt.Sprintf("%d:followCount", in.Userid))
	if err != nil {
		return &pb.GetFollowCountResp{StatusCode: "-1",
				StatusMsg: err.Error(),
				Count:     0},
			errors.Wrapf(errorx.NewDefaultError("redis get err:"+err.Error()), "redis get err ：%v", err)
	}

	if count == "nil" {
		//如果redis中没有则从mysql中拉取并更新至redis中
		db := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
			Where("follower_id = ?", in.Userid).Find(&[]pb.FollowSubject{})
		err = db.Error
		countMysql := db.RowsAffected

		if err != nil {
			return &pb.GetFollowCountResp{StatusCode: "-1",
					StatusMsg: err.Error(),
					Count:     0},
				errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
		}

		l.svcCtx.Rdb.Hset("followCount", fmt.Sprintf("%d:followCount", in.Userid), string(countMysql))

		return &pb.GetFollowCountResp{StatusCode: "0",
			StatusMsg: "get followCount successfully",
			Count:     countMysql}, nil
	}

	countInt, err := strconv.ParseInt(count, 10, 64)

	return &pb.GetFollowCountResp{
		StatusCode: "0",
		StatusMsg:  "get followCount successfully",
		Count:      countInt,
	}, nil

}
