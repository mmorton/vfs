package s3

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/c2fo/vfs/v6"
)

type S3FileOption interface {
	UploadInput(input *s3manager.UploadInput)
}

type ContentTypeOption struct {
	vfs.FileOption

	contentType string
}

var _ S3FileOption = (*ContentTypeOption)(nil)

func WithContentType(contentType string) *ContentTypeOption {
	return &ContentTypeOption{contentType: contentType}
}

func (opt *ContentTypeOption) UploadInput(input *s3manager.UploadInput) {
	input.ContentType = &opt.contentType
}
