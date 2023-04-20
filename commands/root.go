package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-ec2/authenticater"
	"github.com/Appkube-awsx/awsx-ec2/client"
	"github.com/Appkube-awsx/awsx-ec2/commands/ec2cmd"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxEc2Cmd = &cobra.Command{
	Use:   "GetEC2Metadata",
	Short: "GetEC2Metadata command gets resource Arn",
	Long:  `GetEC2Metadata command gets resource Arn details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getEc2Data started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			DescribeInstances(region, acKey, secKey, env, crossAccountRoleArn, externalId)
		}
	},
}

func DescribeInstances(region string, accessKey string, secretKey string, env string, crossAccountRoleArn string, externalId string) *ec2.DescribeInstancesOutput {
	log.Println("Getting aws config resource summary")
	ec2Client := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	ec2Request := &ec2.DescribeInstancesInput{}
	ec2Response, err := ec2Client.DescribeInstances(ec2Request)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	log.Println(ec2Response)
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
	//AwsxEc2Cmd.AddCommand(ec2cmd.GetConfigDataCmd)
	//AwsxEc2Cmd.AddCommand(ec2cmd.GetCostSpikeCmd)
	AwsxEc2Cmd.AddCommand(ec2cmd.GetCostDataCmd)
	AwsxEc2Cmd.AddCommand(ec2cmd.GetCostSpikeCmd)
	AwsxEc2Cmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEc2Cmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEc2Cmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEc2Cmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEc2Cmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEc2Cmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEc2Cmd.PersistentFlags().String("externalId", "", "aws external id auth")
	AwsxEc2Cmd.PersistentFlags().String("env", "", "env")

}

// cmd used to get metadata of EC2 :
//go run main.go --zone=us-east-1 --accessKey=<C6> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
