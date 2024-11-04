package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
)

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
