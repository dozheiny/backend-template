package object

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func GetSpecificPrescriptionPicture(userID uuid.UUID, objectName string, expireTime time.Duration) (*string, error) {

	var generatedLink string

	ctx := context.Background()
	fileAddress := fmt.Sprintf("%s/%s", userID.String(), objectName)

	link, err := minioClient.PresignedGetObject(ctx, prescriptionPicturesBucketName, fileAddress, expireTime, nil)
	if err != nil {
		return nil, err
	}

	generatedLink = link.String()

	if strings.Contains(generatedLink, "172.18.218.103") {
		generatedLink = strings.ReplaceAll(generatedLink, "172.18.218.103", "172.18.4.101")
	}

	if strings.Contains(generatedLink, "172.18.4.103") {
		generatedLink = strings.ReplaceAll(generatedLink, "172.18.4.103", "172.18.4.101")
	}

	return &generatedLink, nil
}
