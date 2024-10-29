package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer // 這是proto產出go檔後會自動生成的接口

	uc *biz.ContentUsecase
}

// NewAppService new a app service.
func NewAppService(uc *biz.ContentUsecase) *AppService {
	return &AppService{uc: uc}
}