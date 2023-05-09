package commands

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-ec2/client"
	"github.com/Appkube-awsx/awsx-ec2/commands/ec2cmd"
	"github.com/Appkube-awsx/awsx-ec2/input"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxEc2Cmd = &cobra.Command{
	Use:   "GetEC2List",
	Short: "GetEC2List command gets list of EC2 instances for a AWS account",
	Long:  `GetEC2List command gets resource Arn details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getEc2Data started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		accessKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()
		//authFlag :=   authenticater.AuthenticateData(vaultUrl, accountNo, region, accessKey, secKey, crossAccountRoleArn, externalId)
		verifiedInput := input.VerifyInputData(vaultUrl, accountNo, region, accessKey, secKey, crossAccountRoleArn, externalId)
		if verifiedInput {
			ListInstances(region, accessKey, secKey, env, crossAccountRoleArn, externalId)
		}
	},
}

func ListInstances(region string, accessKey string, secretKey string, env string, crossAccountRoleArn string, externalId string) *ec2.DescribeInstancesOutput {
	log.Println("Getting aws config resource summary")
	ec2Client := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

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

	return ec2Response
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
