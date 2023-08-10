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

type DeleteFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFollowLogic {
	return &DeleteFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFollowLogic) DeleteFollow(in *pb.DeleteFollowReq) (*pb.DeleteFollowResp, error) {

	cmd := l.svcCtx.Rdb.SRem(l.ctx, strconv.FormatInt(in.ToUserId, 10), in.UserId)
	if cmd.Err() != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: cmd.Err().Error()},
			errors.Wrapf(errorx.NewDefaultError("redis srem err:"+cmd.Err().Error()), "redis srem err ：%v", cmd.Err())

	}

	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ? AND follower_id = ?", in.ToUserId, in.UserId).
		Delete(&pb.FollowSubject{}).
		Error

	if err != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql delete err:"+err.Error()), "mysql delete err ：%v", err)

	}

	//更新redis数据
	//更新followCount
	db := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("follower_id = ?", in.UserId).Find(&[]pb.FollowSubject{})
	err = db.Error
	countMysql := db.RowsAffected

	if err != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}

	l.svcCtx.Rdb.HSet(l.ctx, "followCount", fmt.Sprintf("%d:followCount", in.UserId), strconv.FormatInt(countMysql, 10))

	//更新followerCount
	db = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ?", in.UserId).Find(&[]pb.FollowSubject{})
	err = db.Error
	countMysql = db.RowsAffected

	if err != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}

	l.svcCtx.Rdb.HSet(l.ctx, "followerCount", fmt.Sprintf("%d:followerCount", in.UserId), strconv.FormatInt(countMysql, 10))

	//更新friendCount
	friend := []pb.FollowSubject{}
	err = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ?", in.UserId).Find(&friend).Error
	if err != nil {
		return &pb.DeleteFollowResp{StatusCode: "-1",
				StatusMsg: err.Error(),
			},
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}
	countMysql = 0
	for _, v := range friend {

		err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
			Where("user_id = ? AND follower_id = ?", v.FollowerId, in.UserId).First(&pb.FollowSubject{}).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return &pb.DeleteFollowResp{StatusCode: "-1",
						StatusMsg: err.Error(),
					},
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			}
		}
		countMysql++
	}
	l.svcCtx.Rdb.HSet(l.ctx, "friendCount", fmt.Sprintf("%d:friendCount", in.UserId), strconv.FormatInt(countMysql, 10))

	return &pb.DeleteFollowResp{
		StatusCode: "0",
		StatusMsg:  "delete follow successfully",
	}, nil
}
