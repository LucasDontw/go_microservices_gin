package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Content is a Content model.
type Content struct {
	Id             int64         `json:"id"`
	Title          string        `json:"title" binding:"required"`
	VideoURL       string        `json:"video_url" binding:"required"`
	Author         string        `json:"author" binding:"required"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"category"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int32         `json:"quality"`
	ApprovalStatus int32         `json:"approval_status"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type ContentRepo interface {
	Create(ctx context.Context, c *Content) (error)
	Update(ctx context.Context, id int64, c *Content) error
	IsExist(ctx context.Context, contentID int64) (bool, error)
	Delete(ctx context.Context, id int64) error
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

// UpdateContent update a Content, and returns the update Content.
func (uc *ContentUsecase) UpdateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("UpdateContent: %v", c)

	return uc.repo.Update(ctx, c.Id, c)
}

// DeleteContent delete a Content.
func (uc *ContentUsecase) DeleteContent(ctx context.Context, id int64) error {
	repo := uc.repo
	ok, err := repo.IsExist(ctx, id)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("内容不存在")
	}

	return repo.Delete(ctx, id)
}
