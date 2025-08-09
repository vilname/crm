package rest

import (
	"api/src/model"
	"api/src/service"
	"api/src/util/constant"
	"api/src/util/helper"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAnswer(ctx *gin.Context) {
	page := helper.GetDefaultPage(ctx.Query("page"))
	limit := helper.GetDefaultLimit(ctx.Query("limit"), constant.DefaultPageLimit)

	serviceAnswer := service.NewAnswerService(ctx)
	result, total, err := serviceAnswer.List(page, limit)

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	var pagePagination model.PaginationAnswer
	pagePagination.GetPaginationAnswer(result, page, limit, total)

	ctx.JSON(http.StatusOK, pagePagination)
}

func GetAnswer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	serviceAnswer := service.NewAnswerService(ctx)
	result, err := serviceAnswer.Get(id)

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateAnswer(ctx *gin.Context) {
	var answer model.Answer

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	if err := json.Unmarshal(body, &answer); err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	serviceAnswer := service.NewAnswerService(ctx)
	if err := serviceAnswer.Create(answer); err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}
}
