package services

import (
	"cms/v2/internal/api/operate"
	// "cms/v2/internal/model"
	// "cms/v2/internal/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContentUpdateReq struct {
	Id             int           `json:"id" binding:"required"`
	Title          string        `json:"title"`
	VideoURL       string        `json:"video_url"`
	Author         string        `json:"author"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"catgory"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int           `json:"quality"`
	ApprovalStatus int           `json:"approval_status"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type ContentUpdateRep struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) ContentUpdate(ctx *gin.Context) {
	var req ContentUpdateReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//////////// 使用grpc版本 開始//////////////////
	rseponse, grpcErr := c.operateAppClient.UpdateContent(ctx, &operate.UpdateContentReq{
		Content: &operate.Content{
			Id:             int64(req.Id),
            Title:          req.Title,
            VideoUrl:       req.VideoURL,
            Author:         req.Author,
            Description:    req.Description,
            Thumbnail:      req.Thumbnail,
            Category:       req.Category,
            Duration:       req.Duration.String(),
            Resolution:     req.Resolution,
            FileSize:       req.FileSize,
            Format:         req.Format,
            Quality:        int32(req.Quality),
            ApprovalStatus: int32(req.ApprovalStatus),
		},
	})

	if grpcErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": grpcErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": rseponse,
	})
	//////////// 使用grpc版本 結束//////////////////



	//////////// 直接訪問DB版本 開始//////////////////
	// contentRepo := repositories.NewContentRepo(c.db)
	// ok, err := contentRepo.IsExist(req.Id)

	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "內容不存在"})
	// 	return
	// }

	// now := time.Now()
	// if err := contentRepo.Update(req.Id, model.ContentDetail{
	// 	Title:          req.Title,
	// 	Description:    req.Description,
	// 	Author:         req.Author,
	// 	VideoURL:       req.VideoURL,
	// 	Thumbnail:      req.Thumbnail,
	// 	Category:        req.Category,
	// 	Duration:       req.Duration,
	// 	Resolution:     req.Resolution,
	// 	FileSize:       req.FileSize,
	// 	Format:         req.Format,
	// 	Quality:        req.Quality,
	// 	ApprovalStatus: req.ApprovalStatus,
	// 	Updated_at:     now,
	// }); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "更新失敗"})

	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code":    0,
	// 	"message": "ok",
	// 	"data": &ContentUpdateRep{
	// 		Message: "更新成功",
	// 	},
	// })
	//////////// 直接訪問DB版本 結束//////////////////
}
