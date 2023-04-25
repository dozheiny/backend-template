package object

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
)

// PutObjectOnProfilePictureBucket will put object in student bucket on minio.
func PutObjectOnProfilePictureBucket(id string, file io.Reader, filename string, filesize int64) (*minio.UploadInfo, error) {
	ctx := context.Background()

	filepath := fmt.Sprintf("%s/%s", id, filename)

	result, err := minioClient.PutObject(ctx, prescriptionPicturesBucketName, filepath, file,
		filesize, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &result, nil
}
