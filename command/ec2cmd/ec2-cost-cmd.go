package ec2cmd

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

var GetCostDataCmd = &cobra.Command{
	Use:   "getEC2CostData",
	Short: "Get cost of EC2 services",
	Long:  `Get cost of EC2 services`,
	Run: func(cmd *cobra.Command, args []string) {
		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			instanceName, _ := cmd.Flags().GetString("instanceName")
			if instanceName != "" {
				getClusterCostDetail(*clientAuth)
			} else {
				log.Fatalln("instanceName not provided. Program exit")
			}
		}
	},
}

func getClusterCostDetail(auth client.Auth) (*costexplorer.GetCostAndUsageOutput, error) {
	log.Println("Getting cost data")
	costClient := client.GetClient(auth, client.COST_EXPLORER).(*costexplorer.CostExplorer)
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2023-02-01"),
			End:   aws.String("2023-03-01"),
		},
		Metrics: []*string{
			// aws.String("USAGE_QUANTITY"),
			aws.String("UNBLENDED_COST"),
			aws.String("BLENDED_COST"),
			aws.String("AMORTIZED_COST"),
			aws.String("NET_AMORTIZED_COST"),
			// aws.String("AMORTIZED_COST_FOR_USAGE"),
			// aws.String("NET_UNBLENDED_COST"),
			// aws.String("NORMALIZED_USAGE_AMOUNT"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("REGION"),
			},
		},
		Granularity: aws.String("DAILY"),
		Filter: &costexplorer.Expression{
			And: []*costexplorer.Expression{
				{
					Dimensions: &costexplorer.DimensionValues{
						Key: aws.String("SERVICE"),
						Values: []*string{
							aws.String("Amazon Elastic Compute Cloud - Compute"),
						},
					},
				},
				{
					Dimensions: &costexplorer.DimensionValues{
						Key: aws.String("RECORD_TYPE"),
						Values: []*string{
							aws.String("Credit"),
						},
					},
				},
			},
		},
	}
	costData, err := costClient.GetCostAndUsage(input)
	if err != nil {
		log.Fatalln("Error: in getting cost data", err)
	}
	log.Println(costData)
	return costData, err
}
func init() {
}
