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
		Catgory:         content .GetCategory(),
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

func parseDurationOrZero(durationStr string) time.Duration {
    duration, err := time.ParseDuration(durationStr)
    if err != nil {
        return 0 // 或回傳一個預設值
    }
    return duration
}