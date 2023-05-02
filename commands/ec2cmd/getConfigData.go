package ec2cmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-ec2/authenticater"
	"github.com/Appkube-awsx/awsx-ec2/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

var GetEC2ConfigCmd = &cobra.Command{
	Use:   "getEC2Config",
	Short: "GetEC2Config command gets resource Arn",
	Long:  `GetEC2Config command gets resource Arn details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getEC2Config started")
		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		env := cmd.Parent().PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()
		//instanceName :=

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			instanceName, _ := cmd.Flags().GetString("instanceName")
			if instanceName != "" {
				DescribeInstances(region, acKey, secKey, env, crossAccountRoleArn, externalId, instanceName)
			} else {
				log.Fatalln("instanceName not provided. Program exit")
			}
		}
	},
}

func DescribeInstances(region string, accessKey string, secretKey string, env string, crossAccountRoleArn string, externalId string, instanceName string) *ec2.DescribeInstancesOutput {
	log.Println("Getting aws config resource summary")
	ec2Client := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	filters := []*ec2.Filter{
		{
			Name:   aws.String("tag:Name"),
			Values: []*string{aws.String(instanceName)},
		},
	}
	ec2Request := &ec2.DescribeInstancesInput{
		Filters: filters,
	}
	ec2Response, err := ec2Client.DescribeInstances(ec2Request)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Println(ec2Response)
	return ec2Response
}

func init() {
	GetEC2ConfigCmd.Flags().StringP("instanceName", "t", "", "Instance name")
	if err := GetEC2ConfigCmd.MarkFlagRequired("instanceName"); err != nil {
		log.Fatalln(err)
	}
}

// The command used to get Config data of AWS EC2 instances is given below :

//  ./awsx-ec2 --zone=us-east-1 --accessKey=<#6f> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<> getEC2Config --instanceName=<>
