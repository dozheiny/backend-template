package object

import (
	"net/textproto"
	"strings"
)

func CheckFileType(headers textproto.MIMEHeader) error {

	contentType := headers.Get("Content-Type")

	// Check if content type is images
	if !strings.HasPrefix(contentType, "images/") {
		return FileIsNotImage
	}

	return nil
}
