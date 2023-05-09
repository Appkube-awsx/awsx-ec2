package input

// import (
// 	// "log"

// 	// "github.com/Appkube-awsx/awsx-ec2/vault"
// 	// "github.com/spf13/cobra"
// )

var (
	vaultUrl            string
	accountId           string
	region              string
	accessKey           string
	secretKey           string
	crossAccountRoleArn string
	externalId          string
)

func verifyInputData(vaultUrl string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) bool {
	return true
}

// func verifyInput(cmd *cobra.Command) bool {

// 	vaultUrl = cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
// 	accountId = cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
// 	region = cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
// 	accessKey = cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
// 	secretKey = cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
// 	crossAccountRoleArn = cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
// 	externalId = cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()
// 	verifiedInput := verifyInputData(vaultUrl, accountId, region, accessKey, secretKey, crossAccountRoleArn, externalId)
// 	return verifiedInput
// }

// func verifyInputData(vaultUrl string, accountNo string, region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) bool {
// 	if vaultUrl != "" && accountNo != "" && region != "" && accessKey != "" && secretKey != "" && crossAccountRoleArn != "" && externalId != "" {
// 		data, err := vault.GetAccountDetails(vaultUrl, accountNo)
// 		if err != nil {
// 			log.Println("Error in calling the account details api. \n", err)
// 			return false
// 		}
// 		if data.AccessKey == "" || data.SecretKey == "" || data.CrossAccountRoleArn == "" {
// 			log.Println("Account details not found.")
// 			return false
// 		}
// 		return true
// 	} else {
// 		if accessKey == "" || secretKey == "" || crossAccountRoleArn == "" || externalId == "" {
// 			log.Println("Access Key / Secret Key / Role Arn or External Id not provided ")
// 			return false
// 		}
// 	}
// }
