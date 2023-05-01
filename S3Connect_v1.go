package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	//"gitlab.corp.wabtec.com/atlas/mapa-go"

	noted "ABC"

	"google.golang.org/protobuf/proto"
)

const (
	// Details for the S3 Connection
	s3BucketName = "S3_BUCKET_NAME"
	s3RegionName = "S3_REGION_NAME"
	PartSize     = "PartSize"
	Concurrency  = "Concurrency"
)

type Provider struct {
	sess   *session.Session
	up     *s3manager.Uploader
	down   *s3manager.Downloader
	bucket string
}

// Initialize sets up the application
func (p *Provider) Initialize() {
	os.Setenv(s3BucketName, "")
	os.Setenv(s3RegionName, "us-east-1")
	region := os.Getenv(s3RegionName)
	if region == "" {
		log.Fatal("aws region must be specified through the " + s3RegionName + " env var")
	}
	bucket := os.Getenv(s3BucketName)
	if bucket == "" {
		log.Fatal("aws bucket must be specified through the " + s3BucketName + " env var")
	}
	p.bucket = bucket

	p.sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	p.up = s3manager.NewUploader(p.sess)
	p.up.PartSize, _ = strconv.ParseInt(os.Getenv(PartSize), 10, 64)
	p.up.Concurrency, _ = strconv.Atoi(os.Getenv(Concurrency))
	p.down = s3manager.NewDownloader(p.sess)

}

// Upload uploads the contents of the Note to S3
func (p *Provider) Upload(note noted.Note) (string, error) {
	if len(note.Data) == 0 {
		return "", errors.New("uploaded data must not be nil")
	}
	blb := mapa.Blob{}
	errUnmarshal := proto.Unmarshal(note.Data, &blb)
	if errUnmarshal != nil {
		return "", errUnmarshal
	}

	result, err := p.up.Upload(&s3manager.UploadInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(note.Id),
		Body:   bytes.NewReader(blb.Data),
	})
	log.Println(result, "=result and err===", err)
	if err != nil {
		return "", err
	}
	log.Printf("uploaded %s to %s\n", note.Id, result.Location)
	return result.Location, nil
}

// Download downloads the id specified in the Note
func (p *Provider) Download(note noted.Note) ([]byte, error) {
	if note.GetId() == "" {
		return []byte{}, errors.New("Id cannot be nil")
	}
	buf := []byte{}
	wab := aws.NewWriteAtBuffer(buf)
	_, err := p.down.Download(wab,
		&s3.GetObjectInput{
			Bucket: aws.String(p.bucket),
			Key:    aws.String(note.Id),
		},
	)
	if err != nil {
		return []byte{}, err
	}

	return wab.Bytes(), nil
}
