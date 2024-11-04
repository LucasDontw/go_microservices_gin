package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
	"time"
)

func (a *AppService) CreateContent(ctx context.Context, req *operate.CreateContentReq) (*operate.CreateContentRep, error) {
	content  := req.GetContent();
	uc := a.uc
	err := uc.CreateContent(ctx, &biz.Content{
		Title:           content .GetTitle(),
		VideoURL:        content .GetVideoUrl(),
		Author:          content .GetAuthor(),
		Description:     content .GetDescription(),
		Thumbnail:       content .GetThumbnail(),
		Category:        content .GetCategory(),
		Duration:        parseDurationOrZero(content.GetDuration()),
		Resolution:      content .GetResolution(),
		FileSize:        content .GetFileSize(),
		Format:          content .GetFormat(),
		Quality:         content .GetQuality(),
		ApprovalStatus:  content .GetApprovalStatus(),
	})

	if err != nil {
		return nil, err
	}

	return &operate.CreateContentRep{}, nil
}

func (a *AppService) UpdateContent(ctx context.Context, req *operate.UpdateContentReq) (*operate.UpdateContentRep, error) {
	content  := req.GetContent();
	uc := a.uc
	err := uc.UpdateContent(ctx, &biz.Content{
		Id:              content .GetId(),
		Title:           content .GetTitle(),
		VideoURL:        content .GetVideoUrl(),
		Author:          content .GetAuthor(),
		Description:     content .GetDescription(),
		Thumbnail:       content .GetThumbnail(),
		Category:        content .GetCategory(),
		Duration:        parseDurationOrZero(content.GetDuration()),
		Resolution:      content .GetResolution(),
		FileSize:        content .GetFileSize(),
		Format:          content .GetFormat(),
		Quality:         content .GetQuality(),
		ApprovalStatus:  content .GetApprovalStatus(),
	})

	if err != nil {
		return nil, err
	}

	return &operate.UpdateContentRep{}, nil
}

func (a *AppService) DeleteContent(ctx context.Context,	req *operate.DeleteContentReq) (*operate.DeleteContentRsp, error) {
	uc := a.uc
	err := uc.DeleteContent(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &operate.DeleteContentRsp{}, nil
}


func parseDurationOrZero(durationStr string) time.Duration {
    duration, err := time.ParseDuration(durationStr)

    if err != nil {
        return 0 // 或回傳一個預設值
    }

    return duration
}