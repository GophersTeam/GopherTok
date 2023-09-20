package logic

import (
	"context"
	"sync"

	"GopherTok/common/errorx"
	"GopherTok/server/video/rpc/types/video"

	"github.com/pkg/errors"

	"GopherTok/server/favor/rpc/internal/svc"
	"GopherTok/server/favor/rpc/types/favor"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoredNumOfUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoredNumOfUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoredNumOfUserLogic {
	return &FavoredNumOfUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoredNumOfUserLogic) FavoredNumOfUser(in *favor.FavoredNumOfUserReq) (*favor.FavoredNumOfUserResp, error) {
	list, err := l.svcCtx.VideoRpc.GetUserVideoIdList(l.ctx, &video.GetUserVideoIdListReq{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "err:%v", err)
	}
	var wg sync.WaitGroup
	var sum int64

	resultChan := make(chan int64, len(list.VideoIdList))

	for _, id := range list.VideoIdList {
		wg.Add(1)
		id := id
		go func(videoID int) {
			defer wg.Done()
			num, _ := l.svcCtx.FavorModel.NumOfFavor(l.ctx, id)
			resultChan <- int64(num)
		}(int(id))
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for num := range resultChan {
		sum += num
	}

	return &favor.FavoredNumOfUserResp{
		FavoredNumOfUser: sum,
	}, nil
}
