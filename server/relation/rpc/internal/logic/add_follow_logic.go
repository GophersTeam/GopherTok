package logic

import (
	"context"
	"fmt"
	"strconv"

	"GopherTok/common/errorx"
	"GopherTok/server/relation/dao"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gorm.io/gorm"

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
	intcmd := l.svcCtx.Rdb.SAdd(l.ctx, fmt.Sprintf("cache:gopherTok:follow:id:%d", in.ToUserId), in.UserId)
	if intcmd.Err() != nil {
		return nil,
			errors.Wrapf(errorx.NewDefaultError("redis set err:"+intcmd.Err().Error()), "redis set err ：%v", intcmd.Err())
	}

	kdMysql, err := jsonx.MarshalToString(&dao.FollowData{
		Method:     "creat",
		UserId:     in.ToUserId,
		FollowerId: in.UserId,
		IsFollow:   true,
	})
	if err != nil {
		logx.Errorf("addFollow json.Marshal error: %v", err)
	}
	if err = l.svcCtx.KqPusherMysqlClient.Push(kdMysql); err != nil {
		logx.Errorf("KafkaPusherMysql.Push kdMysql: %s error: %v", kdMysql, err)
	}

	// 更新redis数据
	// 获取followCount

	db := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("follower_id = ?", in.UserId).Find(&[]pb.FollowSubject{})
	err = db.Error
	followCount := db.RowsAffected

	if err != nil {
		return nil,
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}

	// 获取followerCount
	db = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ?", in.UserId).Find(&[]pb.FollowSubject{})
	err = db.Error
	followerCount := db.RowsAffected

	if err != nil {
		return nil,
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}

	// 获取friendCount
	friend := []pb.FollowSubject{}
	err = l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
		Where("user_id = ?", in.UserId).Find(&friend).Error
	if err != nil {
		return nil,
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}
	var friendCount int64 = 0
	for _, v := range friend {

		err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
			Where("user_id = ? AND follower_id = ?", v.FollowerId, in.UserId).First(&pb.FollowSubject{}).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil,
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			}
		}
		friendCount++
	}

	kdRedis, err := jsonx.MarshalToString(&dao.CountData{
		FollowCountKey:   fmt.Sprintf("%d:followCount", in.UserId),
		FollowCount:      strconv.FormatInt(followCount, 10),
		FollowerCountKey: fmt.Sprintf("%d:followerCount", in.UserId),
		FollowerCount:    strconv.FormatInt(followerCount, 10),
		FriendCountKey:   fmt.Sprintf("%d:friendCount", in.UserId),
		FriendCount:      strconv.FormatInt(friendCount, 10),
	})
	if err != nil {
		logx.Errorf("CountData json.Marshal error: %v", err)
	}
	if err = l.svcCtx.KqPusherRedisClient.Push(kdRedis); err != nil {
		logx.Errorf("KafkaPusherRedis.Push kdRedis: %s error: %v", kdMysql, err)
	}

	return &pb.AddFollowResp{
		StatusCode: 0,
		StatusMsg:  "add follow successfully",
	}, nil
}
