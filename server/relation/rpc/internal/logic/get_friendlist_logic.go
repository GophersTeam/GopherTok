package logic

import (
	"GopherTok/common/errorx"
	"GopherTok/server/chat/rpc/chatrpc"
	"GopherTok/server/user/rpc/types/user"
	"context"
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
		return nil,
			errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
	}

	var friendList []*pb.FriendUser
	var idList []int64
	for _, v := range friend {
		err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
			Where("user_id = ? AND follower_id = ?", v.FollowerId, in.Userid).First(&pb.FollowSubject{}).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil,
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			} else {
				return &pb.GetFriendListResp{StatusCode: 0,
					StatusMsg: "get friendList successfully",
					UserList:  nil}, nil
			}

		}

		use, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:        v.FollowerId,
			CurrentId: in.Userid,
		})
		if err != nil {
			return nil,
				errors.Wrapf(errorx.NewDefaultError("userInfo get err:"+err.Error()), "userInfo get err ：%v", err)
		}
		follow1 := pb.FriendUser{
			Id:              use.Id,
			Name:            use.Name,
			FollowCount:     use.FollowCount,
			FollowerCount:   use.FollowerCount,
			IsFollow:        use.IsFollow,
			Avatar:          use.Avatar,
			BackgroundImage: use.BackgroundImage,
			Signature:       use.Signature,
			TotalFavourited: use.TotalFavorited,
			WorkCount:       use.WorkCount,
			FavouriteCount:  use.FavoriteCount,
		}
		friendList = append(friendList, &follow1)
		idList = append(idList, follow1.Id)
	}

	message, err := l.svcCtx.ChatRpc.MessageChatLast(l.ctx, &chatrpc.MessageChatLastRequest{
		FromUserId:   in.Userid,
		ToUserIdList: idList,
	})
	i := 0
	for _, v := range message.LastMessageList {
		friendList[i].Message = v.Content
		friendList[i].MsgType = v.MsgType
	}

	return &pb.GetFriendListResp{StatusCode: 0,
		StatusMsg: "get friendList successfully",
		UserList:  friendList}, nil
}
