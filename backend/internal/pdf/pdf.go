package pdf

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func MergeFiles(files []*multipart.FileHeader, c *gin.Context) (string, error) {
	var filePaths []string
	for _, f := range files {
		path := "./tmp/" + f.Filename
		err := c.SaveUploadedFile(f, path)
		if err != nil {
			return "", err
		}
		filePaths = append(filePaths, path)
	}

	merged := "./tmp/output.pdf"
	err := api.MergeCreateFile(filePaths, merged, false, model.NewDefaultConfiguration())
	if err != nil {
		fmt.Println("Ошибка при слиянии файлов")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", err
	}

	for _, f := range filePaths {
		err := os.Remove(f)
		if err != nil {
			return "", err
		}
	}

	return merged, nil
}

func AddWaterMark(file *multipart.FileHeader, c *gin.Context) (string, error) {
	watermarked := "./tmp/output.pdf"
	tempFilePath := "./tmp/" + file.Filename

	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		return "", err
	}

	wm, err := pdfcpu.ParseTextWatermarkDetails( //todo разобраться с параметрами, вынести в интерфейс
		"CONFIDENTIAL",
		"points:48, color:red, opacity:0.2, rot:45, pos:center, align:center",
		true,
		types.POINTS,
	)

	if err != nil {
		return "", err
	}

	err = api.AddWatermarksFile(
		tempFilePath,
		watermarked,
		[]string{"1"},
		wm,
		model.NewDefaultConfiguration(),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", err
	}

	err = os.Remove(tempFilePath)
	if err != nil {
		return "", err
	}

	fmt.Println("AddWaterMark")
	return watermarked, nil
}
