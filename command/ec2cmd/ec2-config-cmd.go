package ec2cmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
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
		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			instanceName, _ := cmd.Flags().GetString("instanceName")
			if instanceName != "" {
				describeInstances(instanceName, *clientAuth)
			} else {
				log.Fatalln("instanceName not provided. Program exit")
			}
		}
	},
}

func describeInstances(instanceName string, auth client.Auth) *ec2.DescribeInstancesOutput {
	log.Println("Getting aws config resource summary")
	ec2Client := client.GetClient(auth, client.EC2_CLIENT).(*ec2.EC2)
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
