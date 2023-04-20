package client

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-ec2/authenticater"
	"github.com/Appkube-awsx/awsx-ec2/awssession"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetClient(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) *ec2.EC2 {

	sessionName := "assume_role_session_name"
	return assumeRole(crossAccountRoleArn, sessionName, externalId, accessKey, secretKey, region)
}

func GetCostClient(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) *costexplorer.CostExplorer {
	sessionName := "assume_role_session_name"
	return assumeRoleForCost(crossAccountRoleArn, sessionName, externalId, accessKey, secretKey, region)
}

func GetCloudWatchClient() *cloudwatchlogs.CloudWatchLogs {
	sessionName := "assume_role_session_name"
	return assumeRoleForCloudWatch(authenticater.CrossAccountRoleArn, sessionName, authenticater.ExternalId, authenticater.AcKey, authenticater.SecKey, authenticater.Region)
}

func assumeRole(roleArn string, sessionName string, externalId string, accesskey string, secretKey string, region string) *ec2.EC2 {
	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)

	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
		ExternalId:      aws.String(externalId),
	}

	if externalId != "nil" {
		fmt.Println("Trying to fetch external id to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}

	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	//fmt.Println("Assume role output: ", result)

	awsSession, err := awssession.GetSessionByCreds("us-east-1", *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	ec2Client := ec2.New(awsSession)

	return ec2Client
}

func assumeRoleForCost(roleArn string, sessionName string, externalId string, accesskey string, secretKey string, region string) *costexplorer.CostExplorer {
	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)

	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
	}

	if externalId != "nil" {
		fmt.Println("Trying to fetch external id to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}

	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	//fmt.Println("Assume role output: ", result)

	awsSession, err := awssession.GetSessionByCreds("us-east-1", *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	costClient := costexplorer.New(awsSession)

	return costClient
}

func assumeRoleForCloudWatch(roleArn string, sessionName string, externalId string, accesskey string, secretKey string, region string) *cloudwatchlogs.CloudWatchLogs {
	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)

	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
	}

	if externalId != "nil" {
		fmt.Println("Trying to fetch external id to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}

	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	//fmt.Println("Assume role output: ", result)

	awsSession, err := awssession.GetSessionByCreds("us-east-1", *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	costClient := cloudwatchlogs.New(awsSession)

	return costClient
}
