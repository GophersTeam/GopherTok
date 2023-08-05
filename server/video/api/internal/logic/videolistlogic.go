package logic

import (
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"GopherTok/server/video/rpc/types/video"
	"context"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListReq) (resp *types.VideoListResp, err error) {
	// todo: add your logic here and delete this line
	if req.LatestTime == "" {
		req.LatestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}
	UserVideoListCnt, err := l.svcCtx.VideoRpc.VideoList(l.ctx, &video.VideoListReq{
		LatestTime: req.LatestTime,
	})
	list := UserVideoListCnt.VideoList
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	videoList := make([]*types.VideoInfo, 0) // Assuming VideoList is a struct that matches your needs

	for i := 0; i < len(list); i++ {
		videoItem := &types.VideoInfo{
			ID:       list[i].Id,
			Author:   types.AuthorInfo{},
			Title:    list[i].Title,
			PlayURL:  list[i].PlayUrl,
			CoverURL: list[i].CoverUrl,
		}
		videoList = append(videoList, videoItem)
	}

	return &types.VideoListResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
		VideoList: types.VideoList{List: videoList},
	}, nil
}
