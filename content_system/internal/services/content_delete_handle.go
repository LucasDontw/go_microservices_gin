package services

import (
	// "cms/v2/internal/repositories"
	"cms/v2/internal/api/operate"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentDeleteReq struct {
	Id int `json:"id" binding:"required"`
}

type ContentDeleteRep struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//////////// 使用grpc版本 開始//////////////////
	rseponse, grpcErr := c.operateAppClient.DeleteContent(ctx, &operate.DeleteContentReq{
		Id: int64(req.Id),
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

	// if err := contentRepo.Delete(req.Id); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "刪除失敗"})

	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code":    0,
	// 	"message": "ok",
	// 	"data": &ContentDeleteRep{
	// 		Message: "刪除成功",
	// 	},
	// })
	//////////// 直接訪問DB版本 結束//////////////////
}
