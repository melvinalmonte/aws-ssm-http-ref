package clients

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var awsConfig aws.Config
var onceAwsConfig sync.Once

var ssmClient *ssm.Client
var onceSsmClient sync.Once

func getAwsConfig() aws.Config {
	log.Default().Println("-----Initializing AWS config-----")
	onceAwsConfig.Do(func() {
		var err error
		awsConfig, err = config.LoadDefaultConfig(context.Background())
		if err != nil {
			log.Default().Fatal("Unable to load AWS SDK config: ", err.Error())
		}
	})
	log.Default().Println("-----Returning AWS config-----")
	return awsConfig
}

func GetSSMClient() *ssm.Client {
	onceSsmClient.Do(func() {
		log.Default().Println("-----Initializing SSM client-----")

		awsConfig = getAwsConfig()
		log.Default().Println("-----Creating SSM client-----")
		ssmClient = ssm.NewFromConfig(awsConfig)
	})
	return ssmClient

}
