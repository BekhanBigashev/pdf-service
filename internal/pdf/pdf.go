package pdf

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func MergeFiles(files []*multipart.FileHeader) (string, error) {
	var filePaths []string
	for _, f := range files {
		path := "./tmp/" + f.Filename
		err := saveTempFile(f, path)
		if err != nil {
			return "", err
		}
		filePaths = append(filePaths, path)
	}

	merged := "./tmp/output.pdf"
	err := api.MergeCreateFile(filePaths, merged, false, model.NewDefaultConfiguration())
	if err != nil {
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

func AddWaterMark(file *multipart.FileHeader) (string, error) {
	watermarked := "./tmp/output.pdf"
	tempFilePath := "./tmp/" + file.Filename

	err := saveTempFile(file, tempFilePath)
	if err != nil {
		return "", err
	}

	wm, err := pdfcpu.ParseTextWatermarkDetails(
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
		return "", err
	}

	if err := os.Remove(tempFilePath); err != nil {
		return "", err
	}

	return watermarked, nil
}

func saveTempFile(file *multipart.FileHeader, tempFilePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
