package command

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-ec2/command/ec2cmd"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

// AwsxEc2Cmd represents the base command when called without any subcommands
var AwsxEc2Cmd = &cobra.Command{
	Use:   "getEC2List",
	Short: "getEC2List command gets list of EC2 instances for a AWS account",
	Long:  `getEC2List command gets resource Arn details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getEC2List started")
		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			ListInstances(*clientAuth)
		} else {
			cmd.Help()
			return
		}

	},
}

func ListInstances(auth client.Auth) (*ec2.DescribeInstancesOutput, error) {
	log.Println("Getting aws ec2 instance list")
	ec2Client := client.GetClient(auth, client.EC2_CLIENT).(*ec2.EC2)

	ec2Request := &ec2.DescribeInstancesInput{}
	ec2Response, err := ec2Client.DescribeInstances(ec2Request)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	//log.Println(ec2Response)
	for _, reservation := range ec2Response.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("ID: ", *instance.InstanceId, " name: ", *instance.Tags[0].Value)
		}
	}
	log.Println(ec2Response)
	return ec2Response, err
}

func Execute() {
	err := AwsxEc2Cmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxEc2Cmd.AddCommand(ec2cmd.GetEC2ConfigCmd)
	AwsxEc2Cmd.AddCommand(ec2cmd.GetCostSpikeCmd)
	AwsxEc2Cmd.AddCommand(ec2cmd.GetCostDataCmd)

	AwsxEc2Cmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEc2Cmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEc2Cmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEc2Cmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEc2Cmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEc2Cmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEc2Cmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEc2Cmd.PersistentFlags().String("externalId", "", "aws external id auth")
	AwsxEc2Cmd.PersistentFlags().String("env", "", "env")

}

// cmd used to get list of EC2 instance's :

//  ./awsx-ec2 --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
