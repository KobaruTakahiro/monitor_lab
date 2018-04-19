package main

import (
	"fmt"
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/sts"
)

func PrintAwsError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func HandleRequest(ctx context.Context, request events.CloudWatchEvent) {
	t := time.Now()
	t = t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour) // 0時に調整
	start := t.AddDate(0, 0, -1)
	end := start.Add(time.Minute * 5)
	fmt.Printf("%#v \n", t.String())
	fmt.Printf("%#v \n", start.String())
	fmt.Printf("%#v \n", end.String())

	stsClient := NewStsClient()
	input := sts.AssumeRoleInput{
		RoleSessionName: aws.String("session"),
		// arn role
		RoleArn: aws.String(""),
	}
	out, err := stsClient.SVC.AssumeRole(&input)
	if err != nil {
		PrintAwsError(err)
	}
	fmt.Printf("%#v \n", stsClient)	
	fmt.Printf("%#v\n", out)
	cloudwatchClient := NewCloudWatchClient("ap-northeast-1", out.Credentials)

	// vpnTraffic(cloudwatchClient, start, end)
	ec2CPUUtilization(cloudwatchClient, start, end)
}

func createInputData(metric string, namespace string, statistic string, dimentionName string, dimentionValue string, start time.Time, end time.Time) cloudwatch.GetMetricStatisticsInput {
	input := cloudwatch.GetMetricStatisticsInput{
		StartTime: aws.Time(start),
		EndTime: aws.Time(end),
		MetricName: aws.String(metric),
		Namespace: aws.String(namespace),
		Period: aws.Int64(60),
		Statistics: aws.StringSlice([]string{statistic}),
		Dimensions: []*cloudwatch.Dimension{
			{
				Name: aws.String(dimentionName),
				Value: aws.String(dimentionValue),
			},
		},
	}
	return input
}

func vpnTraffic(cloudwatchClient CloudWatchClient, start time.Time, end time.Time) {
	input := createInputData(
		"TunnelDataOut",
		"AWS/VPN",
		"Sum",
		"VpnId",
		"", // vpn id
		start,
		end,
	)
	out, err := cloudwatchClient.SVC.GetMetricStatistics(&input)
	if err != nil {
			PrintAwsError(err)
	}
	for _, data := range out.Datapoints {
		fmt.Printf("%#v \n", data)
	}
}

func ec2CPUUtilization(cloudwatchClient CloudWatchClient, start time.Time, end time.Time) {
	input := createInputData(
		"CPUUtilization",
		"AWS/EC2",
		"Average",
		"InstanceId",
		"", // instance id
		start,
		end,
	)
	out, err := cloudwatchClient.SVC.GetMetricStatistics(&input)
	if err != nil {
			PrintAwsError(err)
	}
	for _, data := range out.Datapoints {
		fmt.Printf("%#v \n", data)
	}
}


func main() {
	lambda.Start(HandleRequest)
}