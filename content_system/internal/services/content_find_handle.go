package services

import (
	"cms/v2/internal/model"
	// "cms/v2/internal/repositories"
	"cms/v2/internal/api/operate"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentFindReq struct {
	Id       int    `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type ContentFindRep struct {
	Message  string                `json:"message" binding:"required"`
	Contents []model.ContentDetail `json:"contents" binding:"required"`
	Total    int64                 `json:"total" binding:"required"`
}

func (c *CmsApp) ContentFind(ctx *gin.Context) {
	var req ContentFindReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//////////// 使用grpc版本 開始//////////////////
	rseponse, grpcErr := c.operateAppClient.FindContent(ctx, &operate.FindContentReq{
		Id:       int64(req.Id),
        Author:   req.Author,
        Title:    req.Title,
        Page:     int32(req.Page),
        PageSize: int32(req.PageSize),
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
	// contentList, count, err := contentRepo.Get(&repositories.FindParams{
	// 	Id:       req.Id,
	// 	Author:   req.Author,
	// 	Title:    req.Title,
	// 	Page:     req.Page,
	// 	PageSize: req.PageSize,
	// })

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // 為何ContentFindRep不要直接設定成[]*model.ContentDetail，那就可以省掉FOR迴圈 => 因為多人應用時直接返回引用可能會出問題，且FOR迴圈深拷貝後安全性較高
	// contents := make([]model.ContentDetail, 0, len(contentList))
	// for _, content := range contentList {
	// 	contents = append(contents, model.ContentDetail{
	// 		Title:          content.Title,
	// 		Description:    content.Description,
	// 		Author:         content.Author,
	// 		VideoURL:       content.VideoURL,
	// 		Thumbnail:      content.Thumbnail,
	// 		Category:        content.Category,
	// 		Duration:       content.Duration,
	// 		Resolution:     content.Resolution,
	// 		FileSize:       content.FileSize,
	// 		Format:         content.Format,
	// 		Quality:        content.Quality,
	// 		ApprovalStatus: content.ApprovalStatus,
	// 		Created_at:     content.Created_at,
	// 		Updated_at:     content.Updated_at,
	// 	})
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code":    0,
	// 	"message": "ok",
	// 	"data": &ContentFindRep{
	// 		Message:  "取得成功",
	// 		Contents: contents,
	// 		Total:    count,
	// 	},
	// })
	//////////// 直接訪問DB版本 結束//////////////////
}
