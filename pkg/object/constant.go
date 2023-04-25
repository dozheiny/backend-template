package object

import (
	"errors"
	"github.com/minio/minio-go/v7"
)

var minioClient *minio.Client

const (
	useSSL = false

	prescriptionPicturesBucketName = "s3-prescription-pictures"
)

var (
	InternalServerError = errors.New("خطای داخلی سرور")
	FileIsNotImage      = errors.New("فایل ارسال شده تصویر نمیباشد")
)
