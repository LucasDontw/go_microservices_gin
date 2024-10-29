package data

import (
	"content_manage/internal/biz"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}

// NewContentRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  // 內容ID
	Title          string        `gorm:"column:title"`           // 內容標題
	Description    string        `gorm:"column:description"`     // 內容描述
	Author         string        `gorm:"column:author"`          // 作者
	VideoURL       string        `gorm:"column:video_url"`       // 視頻撥放URL
	Thumbnail      string        `gorm:"column:thumbnail"`       // 封面圖URL
	Catgory        string        `gorm:"column:catgory"`         // 內容分類
	Duration       time.Duration `gorm:"column:duration"`        // 內容時長
	Resolution     string        `gorm:"column:resolution"`      // 分辨率 ex:720p
	FileSize       int64         `gorm:"column:file_size"`       // 文件大小
	Format         string        `gorm:"column:format"`          // 文件格式 ex:mp4
	Quality        int32           `gorm:"column:quality"`         // 影片品質 1-高清 2-標清
	ApprovalStatus int32           `gorm:"column:approval_status"` // 審核狀態 1-審核中 2-審核通過 3-審核不通過
	Created_at     time.Time     `gorm:"column:created_at"`
	Updated_at     time.Time     `gorm:"column:updated_at"`
}

func (c ContentDetail) TableName() string { //指定此結構對應到哪個table
	return "cms_account.t_content_details"
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	c.log.Infof("contentRepo create content = %+v" , content)
	detail := ContentDetail{
		Title:       content.Title,
        Description: content.Description,
        Author:      content.Author,
        VideoURL:    content.VideoURL,
        Thumbnail:   content.Thumbnail,
        Catgory:     content.Catgory,
        Duration:    content.Duration,
        Resolution:  content.Resolution,
        FileSize:    content.FileSize,
        Format:      content.Format,
        Quality:     content.Quality,
        ApprovalStatus: content.ApprovalStatus,
	}

	db:= c.data.db

	if err:= db.Create(&detail).Error; err != nil {
		c.log.Errorf("content create error =%v", err)

		return err
	}

	return nil
}