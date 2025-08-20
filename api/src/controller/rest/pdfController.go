package rest

import (
	"api/src/util/helper"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ledongthuc/pdf"
	"github.com/nguyenthenguyen/docx"
	"github.com/tealeg/xlsx"
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
	content := ""
	for _, file := range files {
		uploadedFile, err := file.Open()
		if err != nil {
			helper.ErrorResponseMethod(ctx, err)
			return
		}
		defer uploadedFile.Close()

		contentTypes := file.Header.Get("Content-Type")

		contentTmp := ""
		if contentTypes == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
			contentTmp, _ = docxFile(uploadedFile)
		} else if contentTypes == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
			contentTmp, _ = exelFile(uploadedFile)
		} else if contentTypes == "application/pdf" {
			contentTmp, _ = pdfFile(uploadedFile, file.Size)
		}

		content += contentTmp
	}

	ctx.String(http.StatusOK, string(content))
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

func docxFile(uploadedFile multipart.File) (string, error) {
	// Создаем временный файл
	tempFile, err := os.CreateTemp("", "*.docx")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, uploadedFile); err != nil {
		return "", err
	}
	tempFile.Close()

	doc, err := docx.ReadDocxFile(tempFile.Name())
	if err != nil {
		return "", err
	}
	defer doc.Close()

	content := doc.Editable().GetContent()

	return content, nil
}

func exelFile(uploadedFile multipart.File) (string, error) {
	fileData, err := io.ReadAll(uploadedFile)
	if err != nil {
		return "", err
	}

	xlFile, err := xlsx.OpenBinary(fileData)
	if err != nil {
		return "", err
	}

	var sheets string

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			//var rowData []string
			for _, cell := range row.Cells {
				sheets += cell.String()
				sheets += ";"
			}

			sheets += "\n"
		}
	}

	return sheets, nil
}

func pdfFile(uploadedFile multipart.File, fileSize int64) (string, error) {

	data := make([]byte, fileSize)
	_, err := uploadedFile.Read(data)

	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(data)

	text := ""
	textTmp, err := extractTextFromPDFBytes(reader, fileSize)
	text += textTmp + "\n"

	if err != nil {
		return "", err
	}

	return text, nil
}
