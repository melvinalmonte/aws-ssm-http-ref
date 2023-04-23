package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.uber.org/zap"
)

type SSMPutParameterAPI interface {
	PutParameter(ctx context.Context,
		params *ssm.PutParameterInput,
		optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error)
}

func AddStringParameter(c context.Context, api SSMPutParameterAPI, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	zap.S().Info("-----Adding string parameter-----")
	return api.PutParameter(c, input)
}
