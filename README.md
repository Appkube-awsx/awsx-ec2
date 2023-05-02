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

