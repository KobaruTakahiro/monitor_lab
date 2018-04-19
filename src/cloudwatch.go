package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/sts"
)

type CloudWatchClient struct {
	SVC *cloudwatch.CloudWatch
}

func NewCloudWatchClient(region string, cred *sts.Credentials) CloudWatchClient {
	sess := session.Must(session.NewSession())
	creds := credentials.NewStaticCredentials(*cred.AccessKeyId, *cred.SecretAccessKey, *cred.SessionToken)
	svc := cloudwatch.New(sess, aws.NewConfig().WithRegion(region).WithCredentials(creds))

	return CloudWatchClient{SVC: svc}
}
