package logic

import (
	con "GopherTok/common/consts"
	"GopherTok/server/chat/rpc/chatrpc"
	"GopherTok/server/relation/api/internal/svc"
	"GopherTok/server/relation/api/internal/types"
	"GopherTok/server/relation/rpc/pb"
	"GopherTok/server/user/rpc/types/user"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowRes, err error) {
	userid := l.ctx.Value(con.UserId).(int64)
	exists, err := l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: req.ToUserId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if exists.Exists == false {
		return nil, errors.New("user doesn't exist")
	}

	exists, err = l.svcCtx.UserRpc.UserIsExists(l.ctx, &user.UserIsExistsReq{Id: userid})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if exists.Exists == false {
		return nil, errors.New("user doesn't exist")
	}

	if req.ActionType == 1 {
		isFollow, err := l.svcCtx.RelationRpc.CheckIsFollow(l.ctx, &pb.CheckIsFollowReq{
			UserId:   userid,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}

		if isFollow.IsFollow {
			return nil, errors.New("follow has exist")
		}

		follow, err := l.svcCtx.RelationRpc.AddFollow(l.ctx, &pb.AddFollowReq{
			UserId:   userid,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}

		isFriend, _ := l.svcCtx.RelationRpc.CheckIsFollow(l.ctx, &pb.CheckIsFollowReq{
			UserId:   req.ToUserId,
			ToUserId: userid,
		})

		if isFriend.IsFollow {
			_, _ = l.svcCtx.ChatRpc.MessageAction(l.ctx, &chatrpc.MessageActionRequest{
				FromUserId: userid,
				ToUserId:   req.ToUserId,
				Action:     1,
				Content:    "我们已经是好友啦！",
			})
		}
		return &types.FollowRes{
			StatusCode: follow.StatusCode,
			StatusMsg:  follow.StatusMsg,
		}, nil
	} else if req.ActionType == 2 {
		//进行删除操作
		isFollow, err := l.svcCtx.RelationRpc.CheckIsFollow(l.ctx, &pb.CheckIsFollowReq{
			UserId:   userid,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}

		if !isFollow.IsFollow {
			return nil, errors.New("follow doesn't exist")
		}

		follow, err := l.svcCtx.RelationRpc.DeleteFollow(l.ctx, &pb.DeleteFollowReq{
			UserId:   userid,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "req: %+v", req)
		}

		return &types.FollowRes{
			StatusCode: follow.StatusCode,
			StatusMsg:  follow.StatusMsg,
		}, nil
	} else {
		return nil, errors.New("action_type err")
	}

}
