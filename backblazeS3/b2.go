package backblazeS3

import (
	"backblaze_go/config"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type B2Client struct {
	bucketName string
	s3Client   *s3.S3
}

func NewB2Client(b2Cfg config.B2Config) (B2, error) {
	s3Config := &aws.Config{
		Endpoint:    &b2Cfg.Endpoint,
		Region:      &b2Cfg.Region,
		Credentials: credentials.NewStaticCredentials(b2Cfg.KeyID, b2Cfg.ApplicationKey, ""),
	}

	awsSession := session.Must(session.NewSession(s3Config)) //aws session

	s3Client := s3.New(awsSession) // s3 client

	return &B2Client{
		bucketName: b2Cfg.BucketName,
		s3Client:   s3Client,
	}, nil
}

// uploads file at <filePath> to b2 with name <fileName>
func (b *B2Client) Upload(fileName string, filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("[err][b2][upload] failed to open file", err)
		return false
	}
	defer file.Close()
	uploader := s3manager.NewUploaderWithClient(b.s3Client)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &b.bucketName,
		Key:    &fileName,
		Body:   file,
	})
	if err != nil {
		log.Println("[err][b2][upload] failed to upload file", err)
		return false
	}
	return true

}

// downloads file <fileName> to <writeTo>
func (b *B2Client) Download(fileName string, writeTo string) bool {
	file, err := os.Create(writeTo)
	if err != nil {
		log.Println("[err][b2][download] err creating destination file", err)
		return false
	}
	defer file.Close()

	downloader := s3manager.NewDownloaderWithClient(b.s3Client)
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: &b.bucketName,
		Key:    &fileName,
	})
	if err != nil {
		log.Println("[err][b2][download] failed to download file", err)
		return false
	}
	return true
}

// deletes file in b2 with name <fileName>
func (b *B2Client) Delete(fileName string) bool {
	_, err := b.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &b.bucketName,
		Key:    &fileName,
	})
	if err != nil {
		log.Println("[err][b2] failed to delete object", err)
		return false
	}
	return true
}

// Lists all objects in the bucket
func (b *B2Client) List() {
	objects, err := b.s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: &b.bucketName,
	})
	if err != nil {
		log.Println("[err][b2] failed to list objects", err)
		return
	}
	println(len(objects.Contents))

	for _, obj := range objects.Contents {
		println(*obj.Key)
	}
}
