package main


import (
	"time"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	SVC *s3.S3
	Bucket string
}

func NewS3Client(region string, bukect string) S3Client {
	sess := session.Must(session.NewSession())
	creds := credentials.NewEnvCredentials()
	svc := s3.New(sess, aws.NewConfig().WithCredentials(creds))

	return S3Client{SVC: svc, Bucket: bukect}
}

func (client S3Client) SaveObject(service string, region string, id string, label string, data interface{}, t time.Time) error {
	prefix := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s",
		service,
		region,
		t.Format("2006"),
		t.Format("01"),
		t.Format("02"),
		id,
		label,
		t.Format("20060102150405"),
	)

	input := &s3.PutObjectInput{
		Body: aws.ReadSeekCloser(strings.NewReader(fmt.Sprintf("%#v", data))),
		Bucket: aws.String(client.Bucket),
		Key: aws.String(prefix),
	}

	_, err := client.SVC.PutObject(input)
	return err
}

