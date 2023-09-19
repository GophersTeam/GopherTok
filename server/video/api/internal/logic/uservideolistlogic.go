package logic

import (
	"context"
	"strconv"

	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/server/comment/rpc/commentrpc"
	"GopherTok/server/comment/rpc/pb"
	"GopherTok/server/favor/rpc/favorrpc"
	"GopherTok/server/user/rpc/types/user"
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"GopherTok/server/video/rpc/types/video"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVideoListLogic {
	return &UserVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserVideoListLogic) UserVideoList(req *types.UserVideoListReq) (resp *types.UserVideoListResp, err error) {
	// todo: add your logic here and delete this line
	// 获取user id
	uid, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", uid)
	}
	userId, _ := strconv.ParseInt(req.UserId, 10, 64)
	UserVideoListCnt, err := l.svcCtx.VideoRpc.UserVideoList(l.ctx, &video.UserVideoListReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	list := UserVideoListCnt.VideoList

	userinfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		Id:        userId,
		CurrentId: uid,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var (
		videoList    []*types.VideoInfo
		videoResults = make(chan *types.VideoInfo, len(list))
		errorChannel = make(chan error, len(list))
	)

	for i := 0; i < len(list); i++ {
		go func(i int) {
			var commentCount, favoriteCount, isFavorite interface{}
			err := mr.Finish(func() error {
				commentCount, err = l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &pb.GetCommentCountRequest{
					VideoId: list[i].Id,
				})
				return err
			}, func() error {
				favoriteCount, err = l.svcCtx.FavorRpc.FavorNum(l.ctx, &favorrpc.FavorNumReq{
					VideoId: list[i].Id,
				})
				return err
			}, func() error {
				isFavorite, err = l.svcCtx.FavorRpc.IsFavor(l.ctx, &favorrpc.IsFavorReq{
					UserId:  userinfo.Id,
					VideoId: list[i].Id,
				})
				return err
			})
			if err != nil {
				errorChannel <- err
			}

			videoItem := &types.VideoInfo{
				ID: list[i].Id,
				Author: types.AuthorInfo{
					ID:              userinfo.Id,
					Name:            userinfo.Name,
					FollowCount:     userinfo.FollowCount,
					FollowerCount:   userinfo.FollowerCount,
					IsFollow:        userinfo.IsFollow,
					Avatar:          userinfo.Avatar,
					BackgroundImage: userinfo.BackgroundImage,
					Signature:       userinfo.Signature,
					TotalFavorited:  userinfo.TotalFavorited,
					WorkCount:       userinfo.WorkCount,
					FavoriteCount:   userinfo.FavoriteCount,
				},
				Title:         list[i].Title,
				PlayURL:       list[i].PlayUrl,
				CoverURL:      list[i].CoverUrl,
				FavoriteCount: favoriteCount.(*favorrpc.FavorNumResp).Num,
				CommentCount:  commentCount.(*commentrpc.GetCommentCountResponse).Count,
				IsFavorite:    isFavorite.(*favorrpc.IsFavorResp).IsFavor,
			}

			videoResults <- videoItem
		}(i)
	}

	for i := 0; i < len(list); i++ {
		select {
		case videoItem := <-videoResults:
			videoList = append(videoList, videoItem)
		case err := <-errorChannel:
			return nil, errors.Wrapf(err, "req: %+v", req)

		}
	}

	return &types.UserVideoListResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		VideoList: types.VideoList{List: videoList},
	}, nil
}
