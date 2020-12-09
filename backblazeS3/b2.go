package backblazeS3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type B2 interface {
	Upload(string, string) error
	Download(string, string) error
	Delete(string) error
	List() ([]string, error)
}

type B2Client struct {
	bucketName string
	s3Client   *s3.S3
}

// NewB2Client creates a new b2 client with given configuration
func NewB2Client(endpoint, region, keyId, applicationKey, token, bucketName string) (B2, error) {
	s3Config := &aws.Config{
		Endpoint:    &endpoint,
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(keyId, applicationKey, token),
	}

	awsSession, err := session.NewSession(s3Config) //aws session
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(awsSession) // s3 client

	return &B2Client{
		bucketName: bucketName,
		s3Client:   s3Client,
	}, nil
}

// Upload uploads file at <filePath> to b2 with name <fileName>
func (b *B2Client) Upload(fileName string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("[err][b2][upload] failed to open file: '%s'", err)
	}
	defer file.Close()

	uploader := s3manager.NewUploaderWithClient(b.s3Client)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &b.bucketName,
		Key:    &fileName,
		Body:   file,
	})

	if err != nil {
		return fmt.Errorf("[err][b2][upload] failed to upload file: '%s'", err)
	}
	return nil
}

// Download downloads file <fileName> to <writeTo>
func (b *B2Client) Download(fileName string, writeTo string) error {
	file, err := os.Create(writeTo)
	if err != nil {
		return fmt.Errorf("[err][b2][download] err creating destination file: '%s'", err)
	}
	defer file.Close()

	downloader := s3manager.NewDownloaderWithClient(b.s3Client)
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: &b.bucketName,
		Key:    &fileName,
	})

	if err != nil {
		return fmt.Errorf("[err][b2][download] failed to download file: '%s'", err)
	}
	return nil
}

// Delete deletes file in b2 with name <fileName>
func (b *B2Client) Delete(fileName string) error {
	_, err := b.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &b.bucketName,
		Key:    &fileName,
	})
	return err
}

// List lists all objects in the bucket
func (b *B2Client) List() ([]string, error) {
	objects, err := b.s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: &b.bucketName,
	})
	if err != nil {
		return nil, fmt.Errorf("[err][b2] failed to list objects: '%s", err)
	}

	result := make([]string, 0, len(objects.Contents))
	for _, obj := range objects.Contents {
		result = append(result, *obj.Key)
	}

	return result, nil
}
