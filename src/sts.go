package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

type StsClient struct {
	SVC *sts.STS
}

func NewStsClient() StsClient {
	sess := session.Must(session.NewSession())
	creds := credentials.NewEnvCredentials()
	svc := sts.New(sess, aws.NewConfig().WithCredentials(creds))

	return StsClient{SVC: svc}
}