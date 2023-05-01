package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	// Details for the S3 Connection
	s3BucketName = " "  //"S3_BUCKET_NAME"
	s3RegionName = " " //"S3_REGION_NAME"
	partSize     =  //"PartSize"
	File         = "Message.txt"
	retries      = 2
	Concurrency  = "Concurrency"
)

var (
	s3session *s3.S3
)

func int() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(s3RegionName),
	})))
}

type Provider struct {
	sess   *session.Session
	up     *s3manager.Uploader
	down   *s3manager.Downloader
	bucket string
}

func main() {
	file, _ := os.Open(File)
	defer file.Close()

	stats, _ := file.Stat()
	fileSize := stats.Size()

	buffer := make([]byte, fileSize)
	file.Read(buffer)

	expiryDate := time, Now().AddDate(0, 0, 1)
	createResp, err := s3session.CreateMultipartUpload(&s3.CreateMultipartUpload{
		Bucket:  aws.String(s3BucketName),
		Key:     aws.String(file.Name()),
		Expires: &expiryDate,
	})
	if(err !=nil){
		fmt.Println(err)
		return
	}

	var start, currentSize int
	var remaining = int (fileSize)
	var partNum =1
	var completedPart []*s3.completedPart
	for start =0;remaining !=0 start += partSize {
		if remaining <partSize {
			currentSize = remaining
		} else {
			currentSize = partSize
		}
		completed, err :=Upload (createdResp, buffer[start:start+currentsize],partNum)
		if err !=nil {
			_, err =s3session.AbortMultipartUpload(&s3, AportMultipartUploadInput{
				Bucket: createdResp.Bucket,
				Key: createdResp.Key,
				UploadId : createdResp.UploadId
			})
			if err !=nil {
				fmt.Println(err)
				return
			}
		}
		remaining = currentSize
		fmt.Println("Part %v complete, v% bytes remaining", partNum, remaining)
		partNum++
		completedPart = append(completedPart, completed)
	}
	resp, err :=s3session.CreateMultipartUpload(&s3.CreateMultipartUploadInput {
		Bucket: createResp.Bucket,
		Key: createResp.Key,
		Upload: createResp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload {
			Parts: completedPart,
		},
	})
	if err !=nil{
		fmt.Println(err)
	}else {
		fmt.Println(resp.String())
	}
	}

}
func Upload (resp *s3.CreateMultipartUpload, fileBytes []byte, partNum int)

)
if err !=nil {
	fmt.Println(err)
	if try ==retries {
		retrun nil, err
	}else {
		try++
	}
}

