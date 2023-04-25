package object

import (
	"github.com/dozheiny/backend-template/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Initialize will initialize minio object server.
// if there's an error in initialize, it
func Initialize() error {
	endPoint, err := config.Get("MINIO_ENDPOINT")
	if err != nil {
		return err
	}

	accessID, err := config.Get("MINIO_USERNAME")
	if err != nil {
		return err
	}

	secretID, err := config.Get("MINIO_PASSWORD")
	if err != nil {
		return err
	}

	client, err := minio.New(endPoint,
		&minio.Options{Creds: credentials.
			NewStaticV4(accessID, secretID, ""),
			Secure: useSSL,
		})

	if err != nil {
		return err
	}

	minioClient = client

	// check all buckets are exist or not.
	if err := allBucketExistsCheck(); err != nil {
		return err
	}

	return nil
}
