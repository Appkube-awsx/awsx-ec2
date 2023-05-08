- [What is awsx-ec2](#awsx-ec2)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-ec2
This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to EC2 services , primarily the following API's:

- getMetaData
- getConfigData
- setConfigData 
- CostData
- SLAData
- Compliance Data   
- Cost & SLA & Compliance Trends

This cli collect data from metric / logs / traces of the EC2 services and produce the data in a form that Appkube Platform expects.

This CLI , interacts with other Appkube services like Appkube vault , Appkube cloud CMDB so that it can talk with cloud services as 
well as filter and sort the information in terms of product/env/ services, so that Appkube platform gets the data that it expects from the cli.

# How to write plugin subcommand 
Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build / test / debug  / publish and integrate into the main commmand.

# How to build / Test
            go run main.go
                - Program will print Calling aws-cloudelements on console 

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-cloudelements) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
            awsx-cloudelements --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2

# what it does 
This subcommand implement the following functionalities -
   getElementDetails - It  will get the resource count summary for a given AWS account id and region.

# command input
  --valutURL = URL location of vault - that stores credentials to call API
  --acountId = The AWS account id.
  --zone = AWS region
#  command output
{
        ResourceCounts: [
            {
                Count: 124,
                ResourceType: "AWS::S3::Bucket"
            },
            {
                Count: 121,
                ResourceType: "AWS::Lambda::Function"
            },
            {
                Count: 72,
                ResourceType: "AWS::CloudFormation::Stack"
            },
            {
                Count: 50,
                ResourceType: "AWS::CloudWatch::Alarm"
            }
        ],
        TotalDiscoveredResources: 809
}

# How to run 
  From main awsx command , it is called as follows:
  awsx getElementDetails  --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2
  If you build it locally , you can simply run it as standalone command as 
  awsx-cloudelements --vaultURL=vault.dummy.net --accountId=xxxxxxxxxx --zone=us-west-2






# awsx-ec2
EC2 extension

# AWSX Commands for AWSX-EC2 Cli's :

1. CMD used to get list of EC2 instance's :

./awsx-ec2 --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>

2. CMD used to get Config data (metadata) of AWS EC2 instances :

./awsx-ec2 --zone=us-east-1 --accessKey=<#6f> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<> getEC2Config --instanceName=<>

3. CMD used to get cost data of EC2 instances : 

 ./awsx-ec2 --zone=<> --accessKey=<#HH> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>  --env=<dev>  getCostData

4. CMD used to get Cost Data Spike of EC2 instances :

 ./awsx-ec2 --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>  --env=dev GetCostSpike --granularity=DAILY --startDate=2023-03-01 --endDate=2023-03-10

