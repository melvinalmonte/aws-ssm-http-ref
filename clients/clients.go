package clients

import (
	"context"
	"go.uber.org/zap"
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
	zap.S().Info("-----Initializing AWS config-----")
	onceAwsConfig.Do(func() {
		var err error
		awsConfig, err = config.LoadDefaultConfig(context.Background())
		if err != nil {
			zap.S().Error("Unable to load AWS SDK config: ", err.Error())
		}
	})
	zap.S().Info("-----Returning AWS config-----")
	return awsConfig
}

func GetSSMClient() *ssm.Client {
	onceSsmClient.Do(func() {
		zap.S().Info("-----Initializing SSM client-----")
		awsConfig = getAwsConfig()
		zap.S().Info("-----Creating SSM client-----")
		ssmClient = ssm.NewFromConfig(awsConfig)
	})
	return ssmClient

}
