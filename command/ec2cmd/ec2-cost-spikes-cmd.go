package ec2cmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

var dateLayout = "2006-01-02"
var dateTimeLayout = "2006-01-02T15:04:05Z"

var GetCostSpikeCmd = &cobra.Command{
	Use:   "GetEC2CostSpike",
	Short: "Get ec2 cost Spike",
	Long:  `Retrieve ec2 cost spike data from AWS Cost Explorer`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve value of granularity flag
		granularity, err := cmd.Flags().GetString("granularity")
		if err != nil {
			log.Fatalln("Error: in getting granularity flag value", err)
			return
		}
		// Retireve values of start and end date/time
		startDate, err := cmd.Flags().GetString("startDate")
		if err != nil {
			log.Fatalln("Error: in getting startDate flag value", err)
			return
		}
		endDate, err := cmd.Flags().GetString("endDate")
		if err != nil {
			log.Fatalln("Error: in getting endDate flag value", err)
			return
		}
		//authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {

			// wrapperCostSpike(region, crossAccountRoleArn, acKey, secKey, env, externalId, granularity, startDate, endDate, "Amazon Simple Storage Service")
			// wrapperCostSpike(region, crossAccountRoleArn, acKey, secKey, env, externalId, granularity, startDate, endDate, "Amazon CloudFront")
			// wrapperCostSpike(region, crossAccountRoleArn, acKey, secKey, env, externalId, granularity, startDate, endDate, "Amazon DynamoDB")
			// wrapperCostSpike(region, crossAccountRoleArn, acKey, secKey, env, externalId, granularity, startDate, endDate, "Amazon Lambda")
			// wrapperCostSpike(region, crossAccountRoleArn, acKey, secKey, env, externalId, granularity, startDate, endDate, "Amazon API Gateway")
			wrapperCostSpike(*clientAuth, granularity, startDate, endDate, "Amazon Elastic Compute Cloud - Compute")

		}
	},
}

// Wrapper function to get cost, spike percentage and print them.
func wrapperCostSpike(auth client.Auth, granularity string, startDate string, endDate string, service string) (string, error) {
	costClient := client.GetClient(auth, client.COST_EXPLORER).(*costexplorer.CostExplorer)
	fmt.Println("cost spike for: " + service)
	switch granularity {
	case "DAILY":
		// Call CostSpikes function for the date period
		//layout := "2006-01-02" // layout must be the same as start date format
		startDate, err := time.Parse(dateLayout, startDate)
		if err != nil {
			fmt.Println("Error parsing start date:", err)
			return "", err
		}
		endDate, err := time.Parse(dateLayout, endDate)
		if err != nil {
			fmt.Println("Error parsing end date:", err)
			return "", err
		}

		for d := startDate; d.Before(endDate.AddDate(0, 0, 1)); d = d.AddDate(0, 0, 1) {
			prevDate := d.AddDate(0, 0, -1)
			// fmt.Printf("%s (%s)\n", d.Format("2006-01-02"), prevDate.Format("2006-01-02"))
			CostSpikes(costClient, granularity, prevDate.Format(dateLayout), d.Format(dateLayout), service)
		}
		return "", nil

	case "MONTHLY":
		// Call CostSpikes function for the month period
		//layout := "2006-01-02" // layout must be the same as start date format
		startDate, err := time.Parse(dateLayout, startDate)
		if err != nil {
			fmt.Println("Error parsing start date:", err)
			return "", err
		}
		endDate, err := time.Parse(dateLayout, endDate)
		if err != nil {
			fmt.Println("Error parsing end date:", err)
			return "", err
		}

		for d := startDate; d.Before(endDate.AddDate(0, 1, 0)); d = d.AddDate(0, 1, 0) {
			if d.Equal(endDate) {
				break
			}
			prevDate := d.AddDate(0, -1, 0)
			// fmt.Printf("%s (%s)\n", d.Format("2006-01-02"), prevDate.Format("2006-01-02"))
			CostSpikes(costClient, granularity, prevDate.Format(dateLayout), d.Format(dateLayout), service)
		}
		return "", nil

	case "HOURLY":
		// Call CostSpikes function for the hour period
		//layout := "2006-01-02T15:04:05Z" // layout must be the same as start date format
		startDateTime, err := time.Parse(dateTimeLayout, startDate)
		if err != nil {
			fmt.Println("Error parsing start date:", err)
			return "", err
		}
		endDateTime, err := time.Parse(dateTimeLayout, endDate)
		if err != nil {
			fmt.Println("Error parsing end date:", err)
			return "", err
		}

		for d := startDateTime; d.Before(endDateTime); d = d.Add(time.Hour) {
			prevHour := d.Add(-time.Hour)
			// fmt.Println(prevHour.Format(layout), d.Format(layout))
			CostSpikes(costClient, granularity, prevHour.Format(dateTimeLayout), d.Format(dateTimeLayout), service)
		}

		return "", nil

	default:
		return "", fmt.Errorf("unsupported granularity: %s", granularity)
	}

}

// Function to do the cost comparison.
func CostSpikes(costClient *costexplorer.CostExplorer, granularity string, startDateTime string, endDateTime string, service string) (string, error) {
	// Get cost data for latest time period
	startCostData, err := ServiceCostDetails(costClient, granularity, startDateTime, endDateTime, service)
	if err != nil {
		log.Fatalln("Error: in getting cost data for start date", err)
		return "", err
	}

	var endCost float64

	// Get cost data for previous time period
	switch granularity {
	case "MONTHLY":
		// Get start date and end date for previous time period
		previousStartDateTime, previousEndDateTime, err := generateDatesForMonthlyGranularity(startDateTime, endDateTime)
		if err != nil {
			log.Fatalln("Error: in getting previous time period date", err)
			return "", err
		}
		endCostData, err := ServiceCostDetails(costClient, granularity, previousStartDateTime, previousEndDateTime, service)
		if err != nil {
			log.Fatalln("Error: in getting cost data for end date", err)
			return "", err
		}
		endCost = convertCostDataToFloat(endCostData)

	default:
		endCostData, err := ServiceCostDetails(costClient, granularity, endDateTime, endDateTime, service)
		if err != nil {
			log.Fatalln("Error: in getting cost data for end date", err)
			return "", err
		}
		endCost = convertCostDataToFloat(endCostData)
	}

	// Convert cost data to float and positive
	startCost := convertCostDataToFloat(startCostData)
	// endCost = convertCostDataToFloat(endCostData)

	// Calculate cost difference
	costDifference := endCost - startCost

	// Calculate cost difference percentage
	costDifferencePercentage := (costDifference / startCost) * 100

	// |2/12/2023| 5.95|+3%| --- Print format
	if costDifferencePercentage >= 0 {
		output := fmt.Sprintf("|%s| %f | +%f%% |", endDateTime, endCost, costDifferencePercentage)
		fmt.Println(output)
		return output, nil
	}
	if costDifferencePercentage < 0 {
		output := fmt.Sprintf("|%s| %f | %f%% |", endDateTime, endCost, costDifferencePercentage)
		fmt.Println(output)
		return output, nil
	}

	return "", nil
}

// Function to get cost for a given service for given time period.
func ServiceCostDetails(costClient *costexplorer.CostExplorer, granularity string, startDateTime string, endingDateTime string, service string) (string, error) {
	// costClient := client.GetCostClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	// Get endDateTime from startDateTime for DAILY/WEEKLY/HOURLY
	endDateTime, err := generateEndDateTime(granularity, startDateTime)
	if err != nil {
		log.Fatalln("Error: in generating end date time", err)
		return "", err
	}

	var start, end string
	switch granularity {
	case "DAILY":
		// fmt.Println("startDateTime: ", startDateTime, " endDateTime: ", endDateTime)
		start = startDateTime //"2023-03-01"
		end = endDateTime
	case "MONTHLY":
		// fmt.Println( startDateTime, endDateTime)
		start = startDateTime
		end = endDateTime // Use input dates to the function for monthly granularity
	case "HOURLY":
		start = startDateTime
		end = endDateTime
	default:
		return "", fmt.Errorf("unsupported granularity: %s", granularity)
	}

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
			aws.String("BLENDED_COST"),
			aws.String("AMORTIZED_COST"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("RECORD_TYPE"),
			},
		},
		Granularity: aws.String(granularity),
		Filter: &costexplorer.Expression{
			Dimensions: &costexplorer.DimensionValues{
				Key: aws.String("SERVICE"),
				Values: []*string{
					// aws.String("Amazon Elastic Compute Cloud - Compute"),
					aws.String(service),
				},
			},
		},
	}

	costData, err := costClient.GetCostAndUsage(input)
	if err != nil {
		log.Fatalln("Error: in getting cost data", err)
	}

	// fmt.Println("Cost Data: ", costData)

	// Extract the blended cost from the response (change this to get the cost you want)
	blendedCost := float64(0)
	for _, result := range costData.ResultsByTime {
		for _, group := range result.Groups {
			if metrics := group.Metrics; metrics != nil {
				if blendedCostMetric, ok := metrics["BlendedCost"]; ok && blendedCostMetric != nil && blendedCostMetric.Amount != nil {
					if amount, err := strconv.ParseFloat(*blendedCostMetric.Amount, 64); err == nil {
						blendedCost += math.Abs(amount)
					}
				}
			}
		}
	}

	// log.Println(costData.ResultsByTime)
	return strconv.FormatFloat(blendedCost, 'f', -1, 64), err
}

// Function to generate endDateTime according to granularity
func generateEndDateTime(granularity string, startDateTime string) (string, error) {

	switch granularity {
	case "DAILY":
		//layout := "2006-01-02" // layout must be the same as start date format
		startDate, err := time.Parse(dateLayout, startDateTime)
		if err != nil {
			return "", err
		}
		// Add 1 day to the start date to get the end date
		endDate := startDate.AddDate(0, 0, 1)
		end := endDate.Format(dateLayout)

		// fmt.Println("Start Date: ", startDate, "End Date: ", endDate)
		return end, nil

	case "MONTHLY":
		//layout := "2006-01-02" // layout must be the same as start date format
		startDate, err := time.Parse(dateLayout, startDateTime)
		if err != nil {
			return "", err
		}
		// Add 1 month to the start date to get the end date
		endDate := startDate.AddDate(0, 1, 0)
		end := endDate.Format(dateLayout)

		// fmt.Println("Start Date: ", startDate, "End Date: ", endDate)
		return end, nil

	case "HOURLY":
		//layout := "2006-01-02T15:04:05Z" // layout must be the same as start date format
		startDate, err := time.Parse(dateTimeLayout, startDateTime)
		if err != nil {
			return "", err
		}
		// Add 1 hour to the start date to get the end date
		endDate := startDate.Add(time.Hour)
		end := endDate.Format(dateTimeLayout)

		// fmt.Println("Start Date: ", startDate, "End Date: ", endDate)
		return end, nil

	default:
		return "", fmt.Errorf("unsupported granularity: %s", granularity)
	}
}

func generateDatesForMonthlyGranularity(startDateTime string, endingDateTime string) (string, string, error) {
	//layout := "2006-01-02" // layout must be the same as start date format
	startDate, err := time.Parse(dateLayout, startDateTime)
	if err != nil {
		return "", "", err
	}

	endDate, err := time.Parse(dateLayout, endingDateTime)
	if err != nil {
		return "", "", err
	}

	// Calculate time period between start and end date
	timePeriod := endDate.Sub(startDate)

	// Calclulate starting date of previous time period by subtracting time period from start date
	previousStartDate := startDate.AddDate(0, 0, -int(timePeriod.Hours()/24))

	// subtract 1 day from start date to get the end date of previous time period
	previousEndDate := startDate.AddDate(0, 0, -1)

	// Convert the dates to string
	start := previousStartDate.Format(dateLayout)
	end := previousEndDate.Format(dateLayout)

	// fmt.Println("Start Date: ", startDate, "End Date: ", endDate)
	// fmt.Println("Start Date: ", start, "End Date: ", end)
	return start, end, nil
}

func convertCostDataToFloat(CostData string) float64 {
	// Convert the cost data to float
	cost, err := strconv.ParseFloat(CostData, 64)
	if err != nil {
		log.Fatalln("Error: in converting cost data to float", err)
	}

	// Convert the cost to positive if it is negative
	if cost < 0 {
		cost = cost * -1
	}

	return cost
}

func init() {
	GetCostSpikeCmd.Flags().StringP("granularity", "g", "", "granularity")

	if err := GetCostSpikeCmd.MarkFlagRequired("granularity"); err != nil {
		fmt.Println(err)
	}
	GetCostSpikeCmd.Flags().StringP("startDate", "s", "", "startDate")

	if err := GetCostSpikeCmd.MarkFlagRequired("startDate"); err != nil {
		fmt.Println(err)
	}
	GetCostSpikeCmd.Flags().StringP("endDate", "e", "", "endDate")

	if err := GetCostSpikeCmd.MarkFlagRequired("endDate"); err != nil {
		fmt.Println(err)
	}

}

// The command used to get Cost Data Spike of AWS EC2 instances is given below :

//  ./awsx-ec2 --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>  --env=dev GetCostSpike --granularity=DAILY --startDate=2023-03-01 --endDate=2023-03-10
