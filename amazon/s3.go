package amazon

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"

	"github.com/eirka/eirka-libs/config"
)

// Upload a file to S3
func (a *Amazon) Save(filepath, filename, mime string) (err error) {

	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("problem opening file for s3")
	}
	defer file.Close()

	uploader := s3manager.NewUploader(a.session)

	params := &s3manager.UploadInput{
		Bucket:               aws.String(config.Settings.Amazon.Bucket),
		Key:                  aws.String(filename),
		Body:                 file,
		ContentType:          aws.String(mime),
		CacheControl:         aws.String("public, max-age=31536000"),
		ServerSideEncryption: aws.String(s3.ServerSideEncryptionAes256),
	}

	_, err = uploader.Upload(params)

	return

}

// Delete a file from S3
func (a *Amazon) Delete(key string) (err error) {

	svc := s3.New(a.session)

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(config.Settings.Amazon.Bucket),
		Key:    aws.String(key),
	}

	_, err = svc.DeleteObject(params)

	return

}
