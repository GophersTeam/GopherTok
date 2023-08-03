package logic

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/video/api/internal/svc"
	"GopherTok/server/video/api/internal/types"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoReq, r *http.Request) (resp *types.PublishVideoResp, err error) {
	// todo: add your logic here and delete this line

	userId, ok := l.ctx.Value(consts.UserId).(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	// 接收视频流
	file, head, err := r.FormFile("data")
	if err != nil {
		fmt.Printf("Failed to get data, err:%s\n", err.Error())
		return nil, errors.Wrapf(errorx.NewCodeError(40001, errorx.ErrFileOpen), "打开文件错误 err:%v", err)
	}
	defer file.Close()
	// 检查文件类型是否为视频
	if !isVideoFile(head) {
		return nil, errors.Wrapf(errorx.NewDefaultError("文件不是视频类型，请上传视频类型文件"), "文件不是视频类型，请上传视频类型文件 req:%v", req)
	}
	// 计算视频sha256值
	file.Seek(0, 0)
	fileSha256 := utils.FileSha256(file)
	file.Seek(0, 0)
	filePath := "/" + fileSha256 + "/" + head.Filename
	tempFilePath := consts.CoverTemp + fileSha256 + "/" + head.Filename
	err = os.MkdirAll(consts.CoverTemp+fileSha256, 0755)
	defer os.Remove(consts.CoverTemp + fileSha256)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "创建路径 err:%v", err)
	}
	newFile, err := os.Create(tempFilePath)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "创建文件句柄错误 err:%v", err)

	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "io.copy 文件失败 err:%v", err)
	}
	file.Seek(0, 0)

	// Use ffmpeg to extract the first frame as cover
	coverPath := "/" + fileSha256 + "/cover.jpg"
	tempCoverPath := consts.CoverTemp + fileSha256 + "/cover.jpg"
	err = os.MkdirAll(consts.CoverTemp+fileSha256, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "创建路径 err:%v", err)

	}
	defer utils.RemoveContents(consts.CoverTemp + fileSha256)
	ffmpegCmd := exec.Command("ffmpeg", "-i", tempFilePath, "-ss", "00:00:00.001", "-vframes", "1", tempCoverPath)
	ffmpegCmd.Stdin = file
	if err := ffmpegCmd.Run(); err != nil {
		fmt.Println(11111, err)
		return nil, errors.Wrapf(errorx.NewDefaultError("Error while capturing cover"), "Error while capturing cover err:%v", err)
	}

	coverFile, err := os.Open(tempCoverPath)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "open file  err:%v", err)

	}
	defer coverFile.Close()
	var videoUrl, coverUrl string
	switch consts.StoreType(l.svcCtx.Config.CurrentStoreType) {
	case consts.StoreLocal:
	// 本地存储，不做处理不推荐
	case consts.StoreMinio:
		// 使用PutObject上传文件
		videoUrl, err = l.uploadToMinIO(filePath, file)
		if err != nil {
			return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "将video传入minio  err:%v", err)
		}
		coverUrl, err = l.uploadToMinIO(coverPath, coverFile)
		if err != nil {
			return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "将cover传入minio  err:%v", err)
		}
	case consts.StoreCOS:

	}
	fmt.Println(videoUrl)
	fmt.Println("---------------")
	fmt.Println(coverUrl)
	//_, err = l.svcCtx.VideoRpc.PublishVideo(l.ctx, &video.PublishVideoReq{
	//	Id:          l.svcCtx.Snowflake.Generate().Int64(),
	//	UserId:      userId,
	//	Title:       req.Title,
	//	PlayUrl:     videoUrl,
	//	CoverUrl:    coverUrl,
	//	CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
	//	UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
	//	VideoSha256: fileSha256,
	//})
	//if err != nil {
	//	return nil, errors.Wrapf(err, "req: %+v", req)
	//}

	return &types.PublishVideoResp{
		BaseResponse: types.BaseResponse{
			Code:    0,
			Message: "success!",
		},
	}, nil
}

func isVideoFile(fileHeader *multipart.FileHeader) bool {
	// 获取文件的 Content-Type
	contentType := fileHeader.Header.Get("Content-Type")

	// 判断 Content-Type 是否为视频类型
	return strings.HasPrefix(contentType, "video/")
}
func (l *PublishVideoLogic) uploadToMinIO(objectName string, file io.Reader) (string, error) {

	_, err := l.svcCtx.MinioDb.PutObject(l.ctx, consts.MinioBucket, objectName, file, -1, minio.PutObjectOptions{})

	if err != nil {
		logc.Error(l.ctx, err)
		return "", err
	}

	objectURL := l.svcCtx.MinioDb.EndpointURL().String() + "/" + consts.MinioBucket + "/" + objectName
	return objectURL, nil
}
