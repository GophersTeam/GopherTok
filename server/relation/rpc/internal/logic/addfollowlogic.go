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

type AddFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowLogic {
	return &AddFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFollowLogic) AddFollow(in *pb.AddFollowReq) (*pb.AddFollowResp, error) {
	intcmd := l.svcCtx.Rdb.SAdd(l.ctx, strconv.FormatInt(in.ToUserId, 10), in.UserId)
	if intcmd.Err() != nil {
		return &pb.AddFollowResp{StatusCode: "-1",
				StatusMsg: intcmd.Err().Error()},
			errors.Wrapf(errorx.NewDefaultError("redis set err:"+intcmd.Err().Error()), "redis set err ：%v", intcmd.Err())

	}
	err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").Create(&pb.FollowSubject{
		UserId:     in.ToUserId,
		FollowerId: in.UserId,
		IsFollow:   true,
	}).Error
	if err != nil {
		return &pb.AddFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql add err:"+err.Error()), "mysql add err ：%v", err)

	}

	//更新redis数据
	//更新followCount
	db := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("follower_id = ?", in.UserId).Find(&[]pb.FollowSubject{})
	err = db.Error
	countMysql := db.RowsAffected

	if err != nil {
		return &pb.AddFollowResp{StatusCode: "-1",
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
		return &pb.AddFollowResp{StatusCode: "-1",
				StatusMsg: err.Error()},
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}

	l.svcCtx.Rdb.HSet(l.ctx, "followerCount", fmt.Sprintf("%d:followerCount", in.UserId), strconv.FormatInt(countMysql, 10))

	//更新friendCount
	friend := []pb.FollowSubject{}
	err = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ?", in.UserId).Find(&friend).Error
	if err != nil {
		return &pb.AddFollowResp{StatusCode: "-1",
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
				return &pb.AddFollowResp{StatusCode: "-1",
						StatusMsg: err.Error(),
					},
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			}
		}
		countMysql++
	}
	l.svcCtx.Rdb.HSet(l.ctx, "friendCount", fmt.Sprintf("%d:friendCount", in.UserId), strconv.FormatInt(countMysql, 10))

	return &pb.AddFollowResp{StatusCode: "0",
		StatusMsg: "add follow successfully"}, nil
}
