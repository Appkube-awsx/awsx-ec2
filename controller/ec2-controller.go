package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-ec2/command"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func GetEc2ByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*ec2.DescribeInstancesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetEc2sByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetEc2ByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*ec2.DescribeInstancesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetEc2sByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetEc2sByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*ec2.DescribeInstancesOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.ListInstances(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetEc2Instances(clientAuth *client.Auth) (*ec2.DescribeInstancesOutput, error) {
	response, err := command.ListInstances(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
