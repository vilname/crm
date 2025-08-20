package rest

import (
	"api/src/util/helper"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ledongthuc/pdf"
)

type Pdf struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func TextFromPdf(ctx *gin.Context) {
	//file, err := ctx.FormFile("file")

	form, err := ctx.MultipartForm()

	if err != nil {
		helper.ErrorResponseMethod(ctx, err)
		return
	}

	files := form.File["file"]

	text := ""
	for _, file := range files {
		fOpen, err := file.Open()

		if err != nil {
			helper.ErrorResponseMethod(ctx, err)
			return
		}

		defer fOpen.Close()

		data := make([]byte, file.Size)
		_, err = fOpen.Read(data)

		if err != nil {
			helper.ErrorResponseMethod(ctx, err)
			return
		}

		reader := bytes.NewReader(data)

		textTmp, err := extractTextFromPDFBytes(reader, file.Size)
		text += textTmp + "\n"

		if err != nil {
			helper.ErrorResponseMethod(ctx, err)
			return
		}
	}

	ctx.JSON(http.StatusOK, text)
}

func extractTextFromPDFBytes(data io.ReaderAt, size int64) (string, error) {
	reader, err := pdf.NewReader(data, size)
	if err != nil {
		return "", err
	}

	var textBuilder strings.Builder
	totalPage := reader.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := reader.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		// Получаем массив текстовых элементов
		textArray := p.Content().Text

		// Обрабатываем каждый текстовый элемент
		for _, textItem := range textArray {
			textBuilder.WriteString(textItem.S)
		}
	}

	return textBuilder.String(), nil
}
