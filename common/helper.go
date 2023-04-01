package common

import (
	"io"
	"mime/multipart"
	"os"
)

const (
	DefaultLimit    int    = 5
	ImagePath       string = "assets/complaints"
	StatusPending   uint   = 0
	StatusOnReceive uint   = 1
	StatusOnProcess uint   = 2
	StatusFinish    uint   = 3
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func UploadFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func SetStatusFilter(status string) *uint {
	s := new(uint)
	switch status {
	case "0":
		{
			tmp := StatusPending
			s = &tmp
		}
	case "1":
		{
			tmp := StatusOnReceive
			s = &tmp
		}
	case "2":
		{
			tmp := StatusOnProcess
			s = &tmp
		}
	case "3":
		{
			tmp := StatusFinish
			s = &tmp
		}
	default:
	}
	return s
}
