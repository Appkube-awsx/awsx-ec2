package authenticater

import (
	"log"

	"github.com/Appkube-awsx/awsx-ec2/vault"
	"github.com/spf13/cobra"
)

var (
	VaultUrl            string
	AccountId           string
	Region              string
	AcKey               string
	SecKey              string
	CrossAccountRoleArn string
	ExternalId          string
)

func AuthenticateData(vaultUrl string, accountNo string, region string, acKey string, secKey string, crossAccountRoleArn string, externalId string) bool {

	if vaultUrl != "" && accountNo != "" {
		if region == "" {
			log.Fatalln("Zone not provided. Program exit")
			return false
		}
		log.Println("Getting account details")
		data, err := vault.GetAccountDetails(vaultUrl, accountNo)
		if err != nil {
			log.Println("Error in calling the account details api. \n", err)
			return false
		}
		if data.AccessKey == "" || data.SecretKey == "" || data.CrossAccountRoleArn == "" {
			log.Println("Account details not found.")
			return false
		}
		return true

	} else if region != "" && acKey != "" && secKey != "" && crossAccountRoleArn != "" && externalId != "" {
		return true
	} else {
		log.Fatal("AWS credentials like accesskey/secretkey/region/crossAccountRoleArn/externalId not provided. Program exit")
		return false
	}
}

func ChildCommandAuth(cmd *cobra.Command) bool {

	VaultUrl = cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
	AccountId = cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
	Region = cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
	AcKey = cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
	SecKey = cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
	CrossAccountRoleArn = cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
	ExternalId = cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

	authFlag := AuthenticateData(VaultUrl, AccountId, Region, AcKey, SecKey, CrossAccountRoleArn, ExternalId)

	return authFlag
}

// RootCommandAuth -> For validation of parent command
func RootCommandAuth(cmd *cobra.Command) bool {

	VaultUrl = cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
	AccountId = cmd.PersistentFlags().Lookup("accountId").Value.String()
	Region = cmd.PersistentFlags().Lookup("zone").Value.String()
	AcKey = cmd.PersistentFlags().Lookup("accessKey").Value.String()
	SecKey = cmd.PersistentFlags().Lookup("secretKey").Value.String()
	CrossAccountRoleArn = cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
	ExternalId = cmd.PersistentFlags().Lookup("externalId").Value.String()

	authFlag := AuthenticateData(VaultUrl, AccountId, Region, AcKey, SecKey, CrossAccountRoleArn, ExternalId)

	return authFlag
}
