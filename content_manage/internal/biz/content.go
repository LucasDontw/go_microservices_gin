package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Content is a Content model.
type Content struct {
	Title          string        `json:"title" binding:"required"`
	VideoURL       string        `json:"video_url" binding:"required"`
	Author         string        `json:"author" binding:"required"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Catgory        string        `json:"catgory"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int32           `json:"quality"`
	ApprovalStatus int32           `json:"approval_status"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type ContentRepo interface {
	Create(context.Context, *Content) (error)
}

type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

// NewContentUsecase new a Content usecase.
func NewContentUsecase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateContent creates a Content, and returns the new Content.
func (uc *ContentUsecase) CreateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("CreateContent: %v", c)
	return uc.repo.Create(ctx, c)
}